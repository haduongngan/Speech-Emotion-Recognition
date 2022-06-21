# 1. Tổng quan dự án Phân tích cảm xúc cuộc gọi

Đây là dự án bài tập cuối khóa môn Xử lý tiếng nói của nhóm 14. Trong quá trình tìm hiểu lên ý tưởng, nhóm chúng em quan tâm đến bài toán lõi trong xử lý tiếng nói là Phân tích cảm xúc của một đoạn tiếng nói cho trước. Đi từ bài toán lõi đó, chúng em phát triển một phần mềm cho phép phân tích cảm xúc của khách hàng và nhân viên trong các cuộc điện thoại trao đổi, với bối cảnh hướng tới là những trung tâm chăm sóc khách hàng, tổng đài tư vấn,… 

Với ứng dụng này, ngay từ những câu trao đổi đầu tiên, tổng đài có thể phát hiện được cảm xúc của khách hàng, từ đó tự động gợi ý những cách thức giao tiếp hiệu quả. Bên cạnh đó, dựa vào số điện thoại gọi đến, lịch sử cảm xúc những trao đổi của khách hàng với tổng đài trong quá khứ cũng được thống kê. Hai luồng thông tin tham khảo hữu ích đó sẽ giúp nhân viên tổng đài có thể nâng cao chất lượng chăm sóc khách hàng.

Không dừng lại ở đó, ứng dụng cung cấp tính năng phân tích cảm xúc những cuộc gọi đã diễn ra. Cảm xúc của nhân viên và khách hàng trong suốt cuộc gọi đều được xác định và lưu vào cơ sở dữ liệu. Những thông tin này không chỉ phục vụ cho những lần truy vấn sau mà còn giúp nhà quản lý có tư liệu trong đánh giá các đặc trưng nghiệp vụ sau này.

# 2. Mô tả nhóm

Nhóm 14 gồm 4 thành viên: Dương Ngân Hà, Nguyễn Trọng Đạt, Hoàng Hữu Tùng và Đặng Trung Cương. Cụ thể công việc các thành viên đã thực hiện như sau:

- Dương Ngân Hà (19020061): Tìm kiếm ý tưởng ứng dụng, đề xuất luồng hoạt động, tìm hiểu và phụ trách xây dựng huấn luyện model nhận diện cảm xúc tiếng nói.
- Hoàng Hữu Tùng (19020066): Tìm kiếm ý tưởng ứng dụng, tìm kiếm giải pháp cho vấn đề tách giọng nói trong file âm thanh và vấn đề chia nhỏ file âm thanh, xây dựng toàn bộ giao diện ứng dụng, xây dựng router, xử lý call các APIs.
- Nguyễn Trọng Đạt (19021240): Tìm kiếm ý tưởng ứng dụng, phụ trách tính năng xác định giới tính người nói, cài đặt phần sử dụng model đã huấn luyện để nhận diện cảm xúc, xử lý kết quả thu được từ model.
- Đặng Trung Cương (19021229): Tìm kiếm ý tưởng ứng dụng, phụ trách việc nhận diện người nói trong đoạn đối thoại, chia nhỏ file âm thanh, lưu trữ và truy vấn cơ sở dữ liệu.

# 3. Tính năng - Luồng hoạt động của ứng dụng

Trong phần này, nhóm chúng em giới thiệu về các tính năng hệ thống cung cấp và tóm tắt luồng hoạt động của các tính năng đó. 

## 3.1. Tính năng Nhập dữ liệu đầu vào

Demo: [https://drive.google.com/file/d/184IfwOgnXgxgAWPS2lLR5pTb9n0pH9z4/view?usp=sharing](https://drive.google.com/file/d/184IfwOgnXgxgAWPS2lLR5pTb9n0pH9z4/view?usp=sharing)

Với cả tính năng Phân tích lời thoại đầu hay Phân tích cuộc gọi, người dùng đều cần nhập 3 trường dữ liệu: file âm thanh cần phân tích, số điện thoại khách hàng và id của nhân viên tổng đài tiếp nhận cuộc gọi.

- File âm thanh: người dùng có 2 lựa chọn: thu âm trực tiếp hoặc tải file âm thanh sẵn có từ thiết bị của mình. Sau khi hoàn thành việc tải file/thu âm, hệ thống hiển thị âm thanh dạng sóng âm, cho phép người dùng nghe lại, xóa đi nhập lại hoặc ấn save như là cách xác nhận bản ghi này
- Số điện thoại: số điện thoại khách hàng gọi tới tổng đài. Việc cung cấp thông tin này giúp hệ thống có key để truy vấn lịch sử khách hàng trong cơ sở dữ liệu, đồng thời cũng là key lưu trữ để phục vụ truy vấn sau này.
- Id nhân viên: định danh của nhân viên tổng đài - người tiếp nhận cuộc gọi này. Đây là thông tin phục vụ việc lưu trữ trong cơ sở dữ liệu về khách hàng được đầy đủ, đồng thời cũng để quản lý quá trình làm việc của nhân viên.

## 3.2. Luồng tính năng Phân tích lời thoại đầu

Demo: [https://drive.google.com/file/d/1_hrLk5EPhQuS7xkIaQD3cYq6Gm4mUmaa/view?usp=sharing](https://drive.google.com/file/d/1_hrLk5EPhQuS7xkIaQD3cYq6Gm4mUmaa/view?usp=sharing)

Tính năng này cho phép trích xuất các thông tin từ lời thoại đầu tiên của khách hàng. Những thông tin này sẽ giúp nhân viên tiếp nhận cuộc gọi có thêm hiểu biết, từ đó lựa chọn phương thức giao tiếp phù hợp, hiệu quả trong suốt thời gian làm việc với khách hàng sau đó. Với đầu vào là lời thoại đầu tiên của khách hàng khi liên hệ tổng đài, phần mềm sẽ trả về kết quả phân tích cảm xúc (phát hiện cảm xúc của khách hàng khi đến với tổng đài), nhận diện giới tính, và toàn bộ lịch sử cảm xúc trong những lần trao đổi trước đó của khách hàng với tổng đài. 

Chi tiết luồng chạy của tính năng này được trình bày dưới đây: 

- Giả định nghiệp vụ: Lời thoại đầu tiên của khách hàng khi gọi đến tổng đài là trình bày về vấn đề của họ
- Giả định đầu vào: đoạn tiếng nói chỉ có giọng của khách hàng, số điện thoại khách hàng

Với đầu vào như trên, hoạt động của ứng dụng có thể biểu diễn tóm tắt như sơ đồ bên dưới:

[https://drive.google.com/file/d/1hL6jBLoVAC-r2qwPj_2xNhODrgu-P3LE/view?usp=sharing](https://drive.google.com/file/d/1hL6jBLoVAC-r2qwPj_2xNhODrgu-P3LE/view?usp=sharing)

![SP-SER-lời thoại đầu.drawio.png](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F64036976-c618-49b4-8625-11bbe4e12ab6%2FSP-SER-li_thoi_u.drawio.png?table=block&id=d267776c-0d0c-468c-a052-6c11fa3131b8&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)

Với tổ hợp đầu vào, chương trình chia thành 2 nhánh:

- Nhánh truy vấn cơ sở dữ liệu: Dựa vào key là số điện thoại khách hàng, ứng dụng truy vấn toàn bộ dữ liệu cuộc gọi đến từ số điện thoại đó trong lịch sử. Kết quả thao tác (8) trả về lịch sử phân tích cảm xúc khách hàng (9) bao gồm thời điểm diễn ra cuộc gọi, thời lượng, cảm xúc của khách hàng, cảm xúc của nhân viên trong cuộc gọi đó.
- Nhánh sử dụng model để phân tích:
    - Đoạn âm thanh đầu vào có độ dài tùy ý, tuy nhiên vì thiết kế model huấn luyện cùng với tính chất thực tế (cảm xúc của con người có thể biến đổi trong thời gian ngắn), do đó thao tác đầu tiên sau khi ứng dụng nhận được file âm thanh là chia nhỏ chúng thành các đoạn âm ngắn có độ dài 3 giây và thực hiện padding nếu cần (2).
    - Sau khi thu được danh sách các đoạn âm thanh 3 giây, một trong số đó được dùng làm đầu vào cho mô hình nhận diện giới tính (6) để thu được kết quả là giới tính của khách hàng (7)
    - Bên cạnh đó, từng đoạn âm thanh dài 3 giây sẽ đi qua mô hình nhận diện cảm xúc đã huấn luyện (3), đầu ra của thao tác này là danh sách cảm xúc của các đoạn. Trước khi trả về kết quả (5) để hiển thị, danh sách cảm xúc thu được bên trên được xử lý phù hợp (4). Công việc xử lý bao gồm thống kê tỉ lệ các nhãn cảm xúc (số đoạn có cảm xúc đó trên tổng số lượng đoạn âm thanh) và tìm cảm xúc chiếm tỉ lệ lớn nhất (lấy đó làm kết luận về cảm xúc chung của khách hàng).

Các kết quả về phân tích cảm xúc, nhận diện giới tính và lịch sử phân tích sẽ được hiển thị lên giao diện cho người dùng (10).

## 3.3. Luồng tính năng Phân tích cuộc gọi

Ở tính năng này, ứng dụng thực hiện phân tích một cuộc gọi đã hoàn thành và lưu trữ kết quả phân tích vào cơ sở dữ liệu. Những dữ liệu được phân tích và lưu trữ này chính là nguồn tham khảo về khách hàng cho những lần trao đổi kế tiếp. Ngoài ra, hệ thống cũng phân tích cảm xúc của nhân viên tổng đài. Đây có thể coi như một nguồn thông tin giúp các nhà quản lý đánh giá, nâng cao chất lượng dịch vụ.

Chi tiết luồng chạy của tính năng này được trình bày dưới đây: 

- Giả định đầu vào: ghi âm cuộc gọi gồm 2 người tham gia (1 nhân viên tổng đài và 1 khách hàng), số điện thoại khách hàng, id nhân viên tổng đài.

Với đầu vào như trên, hoạt động của ứng dụng có thể biểu diễn tóm tắt như sơ đồ bên dưới

[https://drive.google.com/file/d/1py3xswmXcky_xEy7uEnVXMywrQYZegyg/view?usp=sharing](https://drive.google.com/file/d/1py3xswmXcky_xEy7uEnVXMywrQYZegyg/view?usp=sharing)

![SP-SER-phân tích cuộc gọi.drawio.png](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F5220a822-32ca-4566-a5c7-fc30c2118cab%2FSP-SER-phn_tch_cuc_gi.drawio.png?table=block&id=959a5d16-f6db-4eb8-88e6-e94cf787fe62&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)

Từ đầu vào là một cuộc hội thoại, thao tác đầu tiên ứng dụng thực hiện là nhận diện người nói trong đoạn hội thoại đó (2). Sau thao tác này, đoạn âm thanh nào thuộc về người nói nào sẽ được phát hiện. Để có thể sử dụng model, ứng dụng không chỉ cần phân đoạn theo người nói mà còn cần phân các đoạn có độ dài phù hợp. Do đó, thao tác chia nhỏ (3) cần thực hiện. Những đoạn âm thanh thu được sau quá trình xử lý này đảm bảo rằng mỗi đoạn có độ dài không quá 3 giây và chỉ thuộc về một người nói duy nhất. Chúng em thực hiện phân tích cảm xúc khách hàng và nhân viên theo 2 luồng song song tương tự nhau (chỉ khác đối tượng sở hữu giọng nói) và tương tự quá trình phân tích ở tính năng Phân tích lời thoại đầu, bao gồm các bước:

- Nhận diện giới tính (7),(9) → trả về kết quả giới tính (8), (10)
- Nhận diện cảm xúc của từng đoạn âm thanh nhỏ dài không quá 3 giây của mỗi người (4), (11)
- Xử lý kết quả nhận diện cảm xúc (thống kê tỉ lệ và kết luận về cảm xúc chung) (5), (12)
- Trả về bảng phân tích cảm xúc (6), (13)

Khi đã có tất cả các kết quả nhận diện, phân tích, bên cạnh hiển thị lên giao diện (14) cho người dùng, phần mềm tự động lưu lại thông tin này vào cơ sở dữ liệu (15) phục vụ các tác vụ về sau.

# 4. Dữ liệu

## 4.1. Tập dữ liệu huấn luyện mô hình nhận diện cảm xúc

Để huấn luyện mô hình nhận diện cảm xúc, nhóm chúng em đã sử dụng 2 tập dữ liệu công khai là ****Ryerson Audio-Visual Database of Emotional Speech and Song (RAVDESS) [1]**** và ****Toronto emotional speech set (TESS) [2].****

- ****Ryerson Audio-Visual Database of Emotional Speech and Song (RAVDESS):****
    - Bộ dữ liệu đầy đủ về lời nói và bài hát, âm thanh và video (24,8 GB) có sẵn từ Zenodo. Tuy nhiên với dự án này, chúng em chỉ sử dụng bộ dữ liệu chỉ có tiếng nói (16bit, 48kHz .wav)
    - Bộ dữ liệu chứa 1440 tệp: 60 lần ghi âm cho mỗi người x 24 người = 1440. RAVDESS chứa 24 diễn viên chuyên nghiệp (12 nữ, 12 nam), phát âm hai câu nói ("Kids are talking by the door" và "Dogs are sitting by the door”) bằng giọng Bắc Mỹ trung tính.
    - Cảm xúc lời nói bao gồm các cảm xúc calm, happy, sad, angry, fearful, surprise, and disgust (bình tĩnh, vui vẻ, buồn bã, tức giận, sợ hãi, ngạc nhiên và ghê tởm). Mỗi cảm xúc được tạo ra ở hai mức cường độ cảm xúc (bình thường, mạnh). Ngoài các cảm xúc trên, một phiên bản trạng thái bình thường (neutral) được thêm vào. Người nói sẽ lặp lại 2 lần cho mỗi tổ hợp.
    - Các đoạn âm này đều có độ dài xấp xỉ 3 giây và thường có đoạn đầu im lặng.
- ****Toronto emotional speech set (TESS) [2]:****
    - Đây là tập dữ liệu này chỉ xuất hiện giọng nữ và có chất lượng âm thanh rất cao. Tổng cộng có 2800 tệp âm thanh (định dạng wav) trong bộ dữ liệu này.
    - Với bộ 200 từ vựng có sẵn, 2 nữ diễn viên (26 và 64 tuổi) sẽ ghép chúng vào cấu trúc "Say the word **_”** và thu âm câu hoàn chỉnh (ví dụ: “Say the word bought”).
    - Tổng cộng có 7 cảm xúc được gán nhãn trong bộ dữ liệu này: anger, disgust, fear, happiness, pleasant surprise, sadness, and neutral (tức giận, ghê tởm, sợ hãi, hạnh phúc, ngạc nhiên-thú vị, buồn bã và trung lập).
    - Các đoạn âm thanh này đều có độ dài xấp xỉ 1 giây, và hầu như không có khoảng lặng ở đầu bản ghi.

→ Nhận xét:

- Bọn em sử dụng hai bộ dữ liệu này vì đây là 2 bộ dữ liệu công khai và phổ biến, được đề cập nhiều trong các bài báo bọn em tham khảo. Ngoài ra, nhãn của 2 bộ dữ liệu này có độ trùng khớp cao cũng là yếu tố khiến bọn em quyết định kết hợp
- Khi sử dụng kết hợp 2 bộ dữ liệu này, các nhãn bọn em đồng bộ thành 8 loại `['neutral','calm','happy','sad','angry','fear','disgust','surprise']`
    - Với bộ TESS: “pleasant surprise” → “surprise”, “sadness” → “sad”, “anger” → “angry”, “happiness” → “happy”
    - Với bộ RAVDESS: “fearful” → “fear”
- Ở bộ dữ liệu tổng hợp, số lượng nhãn “calm” ít hơn khá nhiều so với 7 nhãn còn lại.
- Bộ dữ liệu RAVDESS có khoảng lặng ở đầu nhưng bộ dữ liệu TESS thì không. Do đó bọn em cân nhắc việc load dữ liệu thuộc về 2 bộ này một cách riêng biệt, với bộ RAVDESS thì cố tình cắt khoảng 0.5 giây đầu tiên. Tuy nhiên kết quả thử nghiệm cũng không khác biệt so với để toàn bộ (không cắt). Do đó để dễ dàng hơn, trong nhiều thử nghiệm về sau, bọn em đối xử như nhau với tệp âm thanh ở 2 bộ dữ liệu này.

## 4.2. Dữ liệu khi đưa vào sử dụng

Dữ liệu âm thanh đưa vào làm đầu vào của hệ thống là dữ liệu tiếng nói tiếng Anh, định dạng wav. Với từng tính năng đã trình bày, chúng em có giả định về dữ liệu để đảm bảo tính đúng đắn khi chạy hệ thống.

# 5. Phương pháp lõi về âm thanh

Trong phần này, nhóm em trình bày về các phương pháp liên quan đến các tác vụ xử lý âm thanh chính, bao gồm phát hiện người nói trong cuộc hội thoại, nhận diện giới tính và nhận dạng cảm xúc.

## 5.1. Phát hiện người nói trong cuộc hội thoại

Với tính năng phân tích cuộc gọi, đầu vào là đoạn ghi âm có 2 người trao đổi với nhau. Do đó, chúng em cần bộ xử lý để phân tách các đoạn âm thanh nào thuộc về người nói nào. Để thực hiện tác vụ này, nhóm em sử dụng API cung cấp bởi Google (Detect different speakers in an audio recording [3]). Lúc đầu bọn em thử sử dụng phiên bản miễn phí, tuy nhiên nhóm đã gặp nhiều khó khăn do API hoạt động không ổn định. Để ứng dụng hoàn thiện hơn, bọn em sử dụng phiên bản trả phí. 

Khi sử dụng API, bọn em đã viết thêm một số hàm phụ (như convert) nhằm đảm bảo dữ liệu phù hợp định dạng được hỗ trợ và cài đặt thông số hợp lý với nghiệp vụ (số người cần nhận ra là 2, ngôn ngữ là tiếng anh,…)

Đầu vào của API là đoạn âm thanh đã được xử lý cho đúng định dạng, đầu ra API trả về thông tin của từng từ trong đoạn âm đó (thời điểm bắt đầu, thời điểm kết thúc và người nói từ đó). Với đầu ra này, một số thao tác xử lý sẽ được thực hiện để thu được đầu vào thích hợp cho mô hình nhận diện giới tính hoặc nhận diện cảm xúc.

## 5.2. Nhận diện giới tính người nói

Vơi đầu vào là tiếng nói, tiến hành trích xuất đặc trưng trên đoạn âm thanh 2s của người nói đó. Các đặc trưng như MFCC, Mel Spectrogram, Spectrogram, Chromogram, Spectral Contrast.

Các đặc trưng này được đưa qua các mô hình học sâu để nhận diện tiếng nói có thể kể đến như :

### 5.2.1. ResNet50:

Đầu và là các đặc trưng, được đi qua mô hình nhận diện, coi MFCC hay Spectrogram như là đầu vào 1 ảnh để đi qua mô hình ResNet. ResNet là một kiến trúc mạng nơ ron CNN cực kì sâu. Điểm khác biệt của ResNet là nó sử dụng global average pooling layers thay vì sử dụng lớp fully connected layers. Điểm độc đáo nhất của kiến trúc ResNet cố gắng khắc phục vấn đề mất mát thông tin khi ở quá sâu do sử dụng thiết kế nối tắt. Đầu vào $x$ được kết hợp với đầu ra $H(x)$ được xác định bởi :

$$
H\left(x\right)=F\left(x\right)+x
$$

Sự đột phá về kiến trúc này được thể hiện trong hình dưới dây dẫn đến việc ResNet là một mô hình sâu hơn VGGNet tận 8 lần, từ đó độ chính xác cải thiện đáng kể. Ngoài ra ResNet còn sử dụng thêm một số kỹ thuật có thể kể đến như : lớp chuẩn hoá loạt (batch normalization). Chi tiết về kiến trúc mạng ResNet50 được thể hiện trong hình sau

![Untitled](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F65910574-8655-4a98-8040-0f90a4499ebc%2FUntitled.png?table=block&id=42eb1f9a-62c6-46c9-81ff-6b550f68ce72&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)

Kết quả huấn luyện mô hình :

![Untitled](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F03ba1726-b120-4639-9721-1d9b3e094f08%2FUntitled.png?table=block&id=84313148-6826-4db1-9475-fb017b9416df&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)

### 5.2.2 CNN+LSTM+Attention

Kiến trúc mà nhóm em thực nghiệm trong bài toán này được mô tả bởi hình sau đây. Với đầu vào là specspectrogram hoặc MFCC. Dùng CNN để bóc tách thông tin làm đầu vào cho LSTM. Đầu ra ở các trạng thái LSTM được qua Soft Attention để tính xác suất thông tin quan trọng mà muốn đưa qua. Từ những thông tin này, tiến hành đưa qua bộ phân lớp để thu được giới tính của người nói.

![Untitled](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F0c7a8104-b0b5-432c-ae71-15580c8dcb5e%2FUntitled.png?table=block&spaceId=&id=64798dc3-ae14-4cd6-9dc7-4113733183f9&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)

\
\
Kết quả mô hình CNN + LSTM + Attention
![Untitled](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F28ee97d0-945a-470a-9b63-eca4f82aab9e%2FUntitled.png?table=block&id=e202cfd4-b73a-4753-8ece-00694c623579&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)



### 5.2.3 ECAPATDNN

TDNN (Time delay neural network) là một kiến trúc mới tận dụng được ưu điểm là biểu diễn được sự phụ thuộc thời gian dài giữa các đoạn âm thanh, nhưng nhanh hơn rất nhiều so với kiến trúc của RNN hay LSTM do tính toán song song hoá (không bị cản trở tính chất tuần tự của các mạng hồi quy). TDNN sử dụng sub-sampling để giảm bớt tính toán trong quá trình huấn luyện. Kiến trúc của TDNN được thể hiện qua dưới đây

![Untitled](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2Fe68838e3-481c-4c88-bcc7-79768925ffb3%2FUntitled.png?table=block&id=be966f38-32c8-45c1-879e-77a770d94e3b&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)

Phần dưới cùng của kiến trúc TDNN là các đặc trưng của âm thanh được truyền vào. Ở đây tác giả lấy 13 frame trong quá khứ và 9 frame trong tương lai. Trong kiến trúc TDNN, các chuyển đổi ban đầu được học trên các ngữ cảnh hẹp và các lớp sâu hơn sẽ xử lý các đầu ra được qua hàm kích hoạt ẩn từ ngữ cảnh thời gian rộng hơn. Do đó, các tầng cao hơn có khả năng học được các mối quan hệ thời gian rộng hơn. Chúng ta có thể thấy giọng nói hay âm vị có quan hệ với nhau, giả dụ chúng ta có thể dự đoán âm vị ở thời điểm T hiện tại tốt hơn nếu chúng ta có thể biết (học) được đặc trưng bên trái và bên phải một chút so với chỉ sử dụng chính thời điểm T.

Kiến trúc TDNNBase mà nhóm em lựa chọn để thực nghiệm này được tham khảo từ. Đầu ra của x-vector embedding là một vec-tơ có số chiều là 512. Với input đầu vào là ma trận 40xT, với T là tổng số frame của đoạn âm thanh, 40 là dim của đặc trưng MFCC. Ta thu được vector embedding của đoạn âm thanh với 512 chiều. Từ đây đi qua bộ phân lớp và hàm Softmax ta thu được giới tính của người nói.

![Untitled](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F4086a16c-458f-4bed-bdc0-2f9f935b0e66%2FUntitled.png?table=block&id=1134dd1d-fcd5-47b8-be45-25ab46399001&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)

ECAPA-TDNN là kĩ thuật embdeding âm thanh được sử dụng đầu tiên vào bài toán Xác thực tiếng nói. Kiến trúc của ECAPA-TDNN dựa trên cấu trúc x - vector và được cải tiến để tạo ra các embedding tốt hơn. mô tả chi tiết kiến trúc của mạng này. Lớp pooling sử dụng cơ chế chú ý (attention mechanism) phụ thuộc cả vào ngữ cảnh (context) và kênh (channel) cho phép mạng attend vào mỗi khung hình ở các kênh khác nhau. Squeeze and Excitaion (SE) block , do các do các filter hoạt động khai thác thông tin locally, thiếu dữ kiện toàn cảnh nên squeeze sử dụng avg global pooling để thu thập thêm thông tin theo chiều của channel ( chèn thông tin ngữ cảnh toàn cảnh vào các khối cục bộ (locally) . Việc xếp chồng các khối SE-Res2Block nhằm cải thiện hiệu suất của mô hình và giảm kích thước tham số. Cuối cùng, MFA(Multi-layer Feature Aggregation) tổng hợp thông tin toàn bộ layer bằng cách nối các đặc trưng của SE cuối cùng với các đặc trưng SE khác sinh ra bởi các block trước đó.

![Untitled](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2Fe2ff8654-0aaa-412c-8ede-f634dea40680%2FUntitled.png?table=block&id=69a5e59f-0e87-4d03-b644-24a8c7223da3&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)

Mô hình được huấn luyện bằng các tối ưu hoá cực tiểu của hàm AAM softmax. AAM softmax là sự cải tiến vượt bậc so với softmax loss thông thường trong bối cảnh là bài toán xác minh và phân loại chi tiết. Nó trực tiếp tối ưu hoá khoảng các cosine (cosine distance) đo độ giống nhau của 2 vector trong một không gian. AAM softmax là biến thể softmax, ngoài việc phân vùng các vector embedding, nó còn tạo ra các mangular magin để tách biệt các class vector.

Huấn luyện

![Untitled](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F3b14e468-f99f-4720-9217-058c6fd7b80f%2FUntitled.png?table=block&id=f806d596-0d04-4743-8ee2-c0a2718794cc&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)

Kết quả

![Untitled](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F937d9587-0fd6-4372-a254-5b1205affd95%2FUntitled.png?table=block&id=1075c551-9d86-468c-b746-4ff30e1fe49c&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)

Chi tiết về task này được thực hiện tại đây : [datntrong/Speaker-Gender-Recognition-VDT (github.com)](https://github.com/datntrong/Speaker-Gender-Recognition-VDT/tree/main/ipynb_try_model)

## 5.3. Nhận diện cảm xúc của đoạn tiếng nói

Nhóm em đã sử dụng 2 bộ dữ liệu được giới thiệu ở phần trước để huấn luyện mô hình nhận diện cảm xúc của đoạn âm thanh đầu vào. Về cơ bản, việc xây dựng mô hình của nhóm em đi qua các bước sau:

- Làm dày dữ liệu: Do tổng cả 2 bộ dữ liệu trên chỉ khoảng hơn 4000 bản ghi, đồng thời tham khảo trong bài báo [4] cùng một vài nguồn khác, nhóm em thực hiện làm dày dữ liệu bằng cách tạo ra dữ liệu nhiễu (noise speech) và dữ liệu bị kéo dài (stretch and pitch speech). Như vậy, bộ dữ liệu thu được sẽ nhiều lên gấp 3 lần (khoảng gần 13000 bản ghi). Công đoạn này được hỗ trợ bởi thư viện librosa.
- Trích xuất đặc trưng: Từ một đoạn tiếng nói ban đầu, nhóm em tìm cách trích xuất đặc trưng dùng cho huấn luyện và nhận dạng. Nhóm em có tìm hiểu và tham khảo nhiều bài báo, nhất là những bài báo có cùng bộ dữ liệu. Các bài báo [4], [5], [6] đề xuất sử dụng các loại đặc trưng khác nhau và có sự hỗ trợ của các thư viện khác nhau. Do bọn em quyết định sử dụng thư viện librosa nên chỉ chọn thử một số đề xuất. Các đặc trưng bọn em trích xuất bao gồm: zero crossing rate, mfcc, log mel spectrogram, chroma stft, root mean square value, spectral centroid, spectral rolloff.
- Chuẩn hóa đặc trưng: Sau khi các đặc trưng được trích xuất, bọn em thực hiện chuẩn hóa dữ liệu sử dụng thư viện sklearn. Thao tác này xuất phát từ việc quan sát giá trị các đặc trưng được trích xuất, thử nghiệm khi không chuẩn hóa và từ việc tham khảo bài báo [5]. Có 2 kiểu chuẩn hóa nhóm em thử là MinMaxScaler và StandardScaler.
- Chia tập dữ liệu train và test, encode cho label (do đây là bài toán có nhiều nhãn)
- Xây dựng kiến trúc mô hình. Với dự án này, nhóm em đi theo hướng xây dựng các kiến trúc deep learning DNN và CNN (do bọn em nhận thấy trong các bài báo gần đây, các kiến trúc này đạt được hiệu quả hơn các mô hình trước đó). Bọn em đã thử cài đặt 3 kiến trúc khác nhau dựa trên sự tham khảo từ các kiến trúc được đề xuất và chọn sử dụng kiến trúc có độ chính xác cao nhất trên tập test để đưa vào hệ thống.
- Huấn luyện và lưu lại mô hình đã huấn luyện.
- Viết hàm sử dụng mô hình đã huấn luyện để nhận dạng đoạn âm thanh mới.

Với các bước cơ bản như trên, bọn em đã thử các kiến trúc sau:

### 5.3.1. Kiến trúc 1:

- Đặc trưng: ZCR, MFCC, Log Mel-Spectrogram, Chroma, Root Mean Square Value
- Scaler: Standard
- Kiến trúc CNN:
    
    ![model_162.png](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F37a94677-b73d-4711-a874-f6dc58703161%2Fmodel_162.png?table=block&id=db6ab99c-93f4-4497-8567-1a3fa2419b77&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)
    
- Các thông số khác: Batch-size = 64, Learning-rate = 1e-3, loss-function: Cross Entropy Loss, optimizer: Adam

→ Kết quả nhận được khi huấn luyện 150 epochs:

![kt1.png](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F3d29264b-2ff0-43fb-87e7-89b5279c113d%2Fkt1.png?table=block&id=ff276057-a475-4d81-be61-096f06602f46&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)

Đây cũng chính là kiến trúc có độ chính xác cao nhất trên tập test, được sử dụng trong ứng dụng cuối cùng.

### 5.3.2. Kiến trúc 2:

Kiến trúc này nhóm em cài đặt theo một trong 2 phương pháp tiếp cận được đề xuất trong bài báo [5]. Tuy nhiên một số thông số bài báo không đề cập rõ, chúng em đã thử thay đổi và dưới đây là một trong những kết quả chạy tốt nhất theo hướng cài đặt này.

- Đặc trưng: MFCC, Log Mel-Spectrogram, Chroma, Spectral centroid, Spectral rolloff
- Scaler: Standard (bài báo đề xuất MinMax tuy nhiên bọn em thử và thấy không hiệu quả bằng)
- Lưu ý rằng ở bài báo này, nhóm tác giả sử dụng PCA để chọn ra 80 trong 180 đặc trưng. Nhóm em cũng thử nhưng kết quả với thực nghiệm của bọn em cho thấy không sử dụng PCA sẽ tốt hơn, do đó trong tất cả các kiến trúc, bọn em để nguyên số đặc trưng trích xuất được.
- Kiến trúc DNN:
    
    ![kt2.png](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F5be19e17-b615-4081-a414-d6ec6bebfba2%2Fkt2.png?table=block&id=aab24dcd-fcca-4916-93a6-dce14c6baf28&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)
    
- Các thông số khác: Batch-size = 64, Learning-rate = 1e-3, loss-function: Cross Entropy Loss, optimizer: Adam

→ Kết quả nhận được khi huấn luyện 20 epochs:

![kt2_res.png](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F5d5fd13b-1154-4586-a4d6-2289ac81abf3%2Fkt2_res.png?table=block&id=3e13ccd8-f7b8-4e51-9d82-edef1efbc6d5&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)

### 5.3.3. Kiến trúc 3:

Kiến trúc này nhóm em cài đặt theo phương pháp tiếp cận được đề xuất trong bài báo [4]. Tuy nhiên bài báo chỉ đề cập các lớp như thế nào mà không có chi tiết các tham số, chúng em đã thử thay đổi và dưới đây là một trong những kết quả chạy tốt nhất theo hướng cài đặt này.

- Đặc trưng: MFCC, Log Mel-Spectrogram, Chroma
- Scaler: Standard
- Kiến trúc CNN:
    
    ![kt3.png](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F31dbecbc-e8c5-4ad0-805f-9d0b0e0c0ac2%2Fkt3.png?table=block&id=bb269955-7218-4c75-9f64-b8e1a65c7189&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)
    
- Các thông số khác: Batch-size = 64, Learning-rate = 1e-3, loss-function: Cross Entropy Loss, optimizer: Adam

→ Kết quả nhận được khi huấn luyện 15 epochs:

![kt3-res.png](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2Ffc9acd39-7c20-4506-99a6-30155100dad9%2Fkt3-res.png?table=block&id=6b8f205d-153a-4263-beda-f4b60e0d4254&spaceId=c9404b39-5cfc-4fb0-87c5-3c52c8b23828&width=2000&userId=f9c34f43-3a7b-4589-8687-0b440477e789&cache=v2)

Trên đây là 3 bộ tham số và kiến trúc điển hình bọn em thử nghiệm trong quá trình tìm kiếm mô hình hiệu quả nhất. Các lần thử nghiệm khac về cơ bản chỉ thay đổi tham số, tỷ lệ chia dữ liệu, …. nên bọn em không đưa vào đây.

# 6. Tổng kết

## 6.1. Kết luận

Trong dự án này, nhóm em đã cố gắng ứng dụng một bài toán lõi về xử lý tiếng nói để xây dựng một ứng dụng hữu ích, tuy rằng các model chưa đạt được độ chính xác quá cao, luồng chạy chưa tối ưu và gọn gàng. Dựa vào tìm hiểu của nhóm, chúng em tin rằng những hệ thống tương tự nhưng hoàn thiện hơn sẽ có nhiều ứng dụng tuyệt vời trong ngành dịch vụ, không dừng lại ở tổng đài tư vấn mà còn là trung tâm dạy học online, chăm sóc sức khỏe cho trẻ tự kỷ,…

## 6.2. Công nghệ sử dụng

- Giao diện:
    - Sử dụng ReactJS cùng farmework MUI
    - Npm: “wavesurfer" - Hiển thị sóng âm thanh
    - Npm: "react-material-file-upload" -  Upload file
    - Npm: "material-ui-dropzone" - Kéo thả file trên giao diện
- Model nhận diện cảm xúc:
    - Ngôn ngữ: python, thực hiện huấn luyện bằng GPU
    - Package librosa: load dữ liệu, trích xuất đặc trưng
    - Package PyTorch: xây dựng model và huấn luyện
- Backend:
    - Sử dụng go-chi để router API
    - Heroku và Docker để triển khai API lên remote
    - Sử dụng cơ sở dữ liệu PostgreSQL và gorm để tương tác giữa model trong Go và CSDL
    - Sử dụng jwt để quản lý bảo mật
    - Sử dụng swagger để có thể test các API trên cả local và remote
    - Sử dụng Torch api để kết nối với model
    - Sử dụng Flask API

## 6.3. Cấu trúc Github

- Thư mục Backend: chứa các thư mục và tập tin phục vụ cho việc lưu trữ dữ liệu cuộc gọi, khách hàng, và nhân viên nghe máy:
    - Thư mục infrastructure: định nghĩa và load các biến môi trường cho phía backend, phụ trách việc kết nối với database.
    - Thư mục model: định nghĩa các kiểu dữ liệu dùng trong hệ thống như Call (cuộc gọi), Customer (khách hàng), Staff (nhân viên tổng đài), đồng thời định nghĩa các hàm sẽ được sử dụng.
    - Thư mục repository: Cài đặt các hàm đã được định nghĩa trong thư mục model, chủ yếu là để tương tác với cơ sở dữ liệu (thêm, sửa và xóa dữ liệu).
    - Thư mục service: Định nghĩa và cài đặt các hàm xử lý nghiệp vụ cho hệ thống (nếu có)
    - Thư mục controller: Định nghĩa và cài đặt các API cho hệ thống.
    - Thư mục router: Định tuyến các API để phía frontend có thể gọi.
    - Dockerfile, heroku.yml: đóng gói các quy trình để có thể triển khai hệ thống lên máy chủ, cho phép truy cập từ xa.
    - Flask API : API trả về cho frontend nhận diện cảm xúc, nhận diện giới tính người nói. Các component nhằm kết nối được với các API của google và các model cảm xúc, nhận diện giới tính người nói.
- Thư mục FrontEnd: chứa thư mục và tệp tin toàn bộ giao diện hệ thống với các chức năng cơ bản là phân tích lời thoại đầu, phân tích cuộc gọi và xem lịch sử phân tích.
    - Thư mục src/component: chứa các component dùng chung
    - Thư mục src/compoent/table: chứa các bảng phân tích dữ liệu và lich sử
    - Thư mục src/page: chứa các trang hiển thị giao diện
    - Thư mục rc/apis: chứa các xử lý call APIs
- Thư mục Model_Emo: chứa code cài đặt xây dựng và huấn luyện mô hình, mô hình đã huấn luyện, scaler tương ứng:
    - Các tệp có tên Model-x.ipynb: code cài đặt huấn luyện model lần lượt theo các kiến trúc 1, 2 và 3
    - Các tệp có tên model_x.pkl: mô hình đã huấn luyện rồi được lưu lại sử dụng cho nhận dạng đoạn âm thanh mới
    - Các tệp có tên scaler_x.pkl: scaler được sử dụng tương ứng với từng mô hình đã huấn luyện, sử dụng khi nhận dạng đoạn âm thanh mới.

# Tài liệu tham khảo - Liên kết

[1] Bộ dữ liệu RAVDESS :[https://www.kaggle.com/datasets/uwrfkaggler/ravdess-emotional-speech-audio](https://www.kaggle.com/datasets/uwrfkaggler/ravdess-emotional-speech-audio)

[2] Bộ dữ liệu TESS: [https://tspace.library.utoronto.ca/handle/1807/24487](https://tspace.library.utoronto.ca/handle/1807/24487)

[3] API Detect different speakers in an audio recording: [https://cloud.google.com/speech-to-text/docs/multiple-voices#speech_transcribe_diarization_beta-protocol](https://cloud.google.com/speech-to-text/docs/multiple-voices#speech_transcribe_diarization_beta-protocol)

[4] Detection of Emotion of Speech for RAVDESS Audio Using Hybrid Convolution Neural Network: [https://www.hindawi.com/journals/jhe/2022/8472947/](https://www.hindawi.com/journals/jhe/2022/8472947/) → [https://www.kaggle.com/dngngnh/sp-ser-hybrid](https://www.kaggle.com/dngngnh/sp-ser-hybrid)

[5] Two-Way Feature Extraction for Speech Emotion Recognition Using Deep Learning: [https://pubmed.ncbi.nlm.nih.gov/35336548/](https://pubmed.ncbi.nlm.nih.gov/35336548/)

[6] Ensemble Learning of Hybrid Acoustic Features for Speech Emotion Recognition: [https://www.researchgate.net/publication/340097289_Ensemble_Learning_of_Hybrid_Acoustic_Features_for_Speech_Emotion_Recognition](https://www.researchgate.net/publication/340097289_Ensemble_Learning_of_Hybrid_Acoustic_Features_for_Speech_Emotion_Recognition)

[7] A Real-Time Speech Emotion Recognition System and its Application in Online Learning: [https://www.sciencedirect.com/science/article/pii/B9780128018569000025](https://www.sciencedirect.com/science/article/pii/B9780128018569000025)