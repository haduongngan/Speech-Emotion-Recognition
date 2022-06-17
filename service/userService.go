package service

import (
	"errors"
	"log"
	"spser/infrastructure"
	"spser/middleware"
	"spser/model"
	"spser/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository model.UserRepository
}

type UserService interface {
	GetAll() ([]model.User, error)
	CreateUser(newUser *model.User) (*model.UserResponse, error)
	GetById(id int) (*model.UserResponse, error)
	GetByUsername(username string) (*model.UserResponse, error)
	DeleteUser(id int) (*model.User, error)
	CheckCredentials(id int, password string) (*model.User, error)
	LoginRequest(username string, password string) (*model.User, string, string, error)
	LoginWithToken(token string) (*model.User, string, string, bool, error)
}

func (s *userService) GetAll() ([]model.User, error) {
	return s.userRepository.GetAll()
}

func (s *userService) CreateUser(newUser *model.User) (*model.UserResponse, error) {
	newUser.Password = hashAndSalt(newUser.Password)
	newUser, err := s.userRepository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	newUserResponse := model.UserResponse{
		Id:          newUser.Id,
		Username:    newUser.Username,
		Role:        newUser.Role,
		CompanyName: newUser.CompanyName,
	}
	return &newUserResponse, nil
}

func (s *userService) GetById(id int) (*model.UserResponse, error) {
	user, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return &model.UserResponse{
		Id:          user.Id,
		Username:    user.Username,
		Role:        user.Role,
		CompanyName: user.CompanyName,
	}, nil

}

func (s *userService) DeleteUser(id int) (*model.User, error) {
	return s.userRepository.DeleteUser(id)
}

func (s *userService) LoginRequest(username string, password string) (*model.User, string, string, error) {
	// validate username/password
	user, err := s.userRepository.GetByUsername(username)
	if err != nil {
		return nil, "", "", err
	}

	err = checkPassword(user, password)
	if err != nil {
		return nil, "", "", err
	}

	// get JWT
	tokenString, refreshToken, err := middleware.GetTokenString(user)
	if err != nil {
		infrastructure.ErrLog.Printf("Problem with Login Request - error getting JWT: %v,n", err)
		return nil, "", "", err
	}

	// err = s.userRepository.ChangePermissionWhenLogin(username)
	// if err != nil {
	// 	infrastructure.ErrLog.Printf(err.Error())
	// 	return user, tokenString, refreshToken, err
	// }
	return user, tokenString, refreshToken, nil

}

func (s *userService) LoginWithToken(token string) (*model.User, string, string, bool, error) {
	user, err := middleware.GetClaimsData(token)
	if err != nil {
		infrastructure.ErrLog.Println("Problem with LoginWithToken at GetClaimsData: ", err)
		return nil, "Invalid token", "", false, err
	}

	timeLeft := user.VerifyExpiresAt(time.Now().UnixNano()/infrastructure.NANO_TO_SECOND, true)
	if !timeLeft {
		infrastructure.ErrLog.Printf("Session expired!")
		return nil, "Token exceeded expire time!", "", false, nil
	}

	if ok, err := s.userRepository.LoginTokenRequest(user); err != nil {
		infrastructure.ErrLog.Printf("Problem with LoginWithToken: %v/n", err)
	} else {
		if !ok {
			return nil, "Invalid token", "", false, nil
		}
	}

	newTokenString, newRefreshTokenString, err := middleware.GetTokenString(user)
	if err != nil {
		infrastructure.ErrLog.Printf("Problem with LoginRequest at GetTokenString: %v/n", err)
		return nil, "", "", false, err
	}

	return user, newTokenString, newRefreshTokenString, true, nil
}
func (s *userService) GetByUsername(username string) (*model.UserResponse, error) {
	user, err := s.userRepository.GetByUsername(username)
	if err != nil {
		infrastructure.ErrLog.Printf(err.Error())
		return nil, err
	}

	userResponse := &model.UserResponse{
		Id:          user.Id,
		Role:        user.Role,
		Username:    user.Username,
		CompanyName: user.CompanyName,
	}
	return userResponse, nil
}
func NewUserService() UserService {
	return &userService{
		userRepository: repository.NewUserRepository(),
	}
}

func checkPassword(user *model.User, password string) error {
	// check password: hashed one.
	// compare: hashed, plain
	if !comparePassword(user.Password, password) {
		return errors.New("incorrect password from service/checkPassword")
	}

	return nil
}

func (s *userService) CheckCredentials(id int, password string) (*model.User, error) {
	user, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	if !comparePassword(user.Password, password) {
		return nil, errors.New("incorrect password from service/check credential")
	}
	return user, nil
}
func comparePassword(hashedPwd string, plainPwd string) bool {
	// cmphashandpass : true ~ nil.
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd)); err != nil {
		return false
	}
	return true
}

func hashAndSalt(password string) string {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err.Error() + "/service/hashAndSalt")
	}
	// log.Println("hashedPwd is:", string(hashedPwd))
	return string(hashedPwd)
}
