
import time

import torch
device = torch.device('cuda:0' if torch.cuda.is_available() else 'cpu')

import os
from torch.utils.data import Dataset, DataLoader
import pandas as pd
import numpy as np
import random
from time import time
from collections import Counter
from scipy.io import wavfile
from scipy import signal
import scipy
from PIL import Image


def log_specgram(audio, sample_rate, window_size=20, step_size=10, eps=1e-10):
    nperseg = int(round(window_size * sample_rate / 1e3))
    noverlap = int(round(step_size * sample_rate / 1e3))
    freqs, _, spec = signal.spectrogram(audio,
                                        fs=sample_rate,
                                        window='hann',  # 'text'
                                        nperseg=nperseg,
                                        noverlap=noverlap,
                                        detrend=False)
    return np.log(spec.T.astype(np.float32) + eps)


def random_segment(audio_signal, N):
    length = audio_signal.shape[0]
    if N < length:
        start = random.randint(0, length - N)
        audio_signal = audio_signal[start:start + N]
    else:
        tmp = np.zeros((N,))
        start = random.randint(0, N - length)
        tmp[start: start + length] = audio_signal
        audio_signal = tmp
        # test_sound = np.pad(test_sound, (N - test_sound.shape[0])//2, mode = 'constant')
    return audio_signal


def gen_spec(wav_path, duration):
    samplerate, test_sound = wavfile.read(wav_path)
    test_sound = test_sound
    N = int(duration * samplerate)
    segment_sound = random_segment(test_sound, N)
    spectrogram = log_specgram(segment_sound, samplerate).astype(np.float32)

    spectrogram = np.array(Image.fromarray(spectrogram).resize(size=(300, 300)))

    out = np.zeros((3, 300, 300), dtype=np.float32)
    out[0, :, :] = spectrogram
    out[1, :, :] = spectrogram
    out[2, :, :] = spectrogram
    return out


class MyDatasetSTFT(Dataset):
    def __init__(self, filenames, labels, transform=None, duration=2, test=False):
        assert len(filenames) == len(labels), "Number of files != number of labels"
        self.fns = filenames
        self.lbs = labels
        self.transform = transform
        self.duration = duration
        self.test = test

    def __len__(self):
        return len(self.fns)

    def __getitem__(self, idx):
        if self.test:
            fname = self.fns[idx]

        else:
            fname = self.fns[idx]

        feats = gen_spec(fname, self.duration)
        return feats, self.lbs[idx], self.fns[idx]








# from __future__ import print_function
import os
from torch.autograd import Variable
from torch.utils.data import Dataset
import torchvision.models as models
import torch
import torch.nn as nn
import pickle
import pdb
import torch.optim as optim
from PIL import Image
import numpy as np
import random
import torch.backends.cudnn as cudnn
from time import time
from scipy.io import wavfile



def net_frozen(args, model):
    print('********************************************************')
    model.frozen_until(args.frozen_until)
    init_lr = args.lr
    if args.trainer.lower() == 'adam':
        optimizer = optim.Adam(filter(lambda p: p.requires_grad, model.parameters()),
                lr=init_lr, weight_decay=args.weight_decay)
    elif args.trainer.lower() == 'sgd':
        optimizer = optim.SGD(filter(lambda p: p.requires_grad, model.parameters()),
                lr=init_lr,  weight_decay=args.weight_decay)
    print('********************************************************')
    return model, optimizer

def parallelize_model(model):
    if torch.cuda.is_available():
        model = model.cuda()
        model = torch.nn.DataParallel(model, device_ids=range(torch.cuda.device_count()))
        cudnn.benchmark = True
    return model

def unparallelize_model(model):
    try:
        while 1:
            # to avoid nested dataparallel problem
            model = model.module
    except AttributeError:
        pass
    return model

def second2str(second):
    h = int(second/3600.)
    second -= h*3600.
    m = int(second/60.)
    s = int(second - m*60)
    return "{:d}:{:02d}:{:02d} (s)".format(h, m, s)

def print_eta(t0, cur_iter, total_iter):
    """
    print estimated remaining time
    t0: beginning time
    cur_iter: current iteration
    total_iter: total iterations
    """
    time_so_far = time() - t0
    iter_done = cur_iter + 1
    iter_left = total_iter - cur_iter - 1
    second_left = time_so_far/float(iter_done) * iter_left
    s0 = 'Epoch: '+ str(cur_iter + 1) + '/' + str(total_iter) + ', time so far: ' \
        + second2str(time_so_far) + ', estimated time left: ' + second2str(second_left)
    print(s0)

def get_length_wav(fn):
    frame_rate, signal = wavfile.read(fn)
    return float(signal.shape[0])/frame_rate













# from utils import *
# from data import MyDatasetSTFT
from sklearn.model_selection import train_test_split
# from utils import *
import torch
import torch.nn as nn
import argparse
import copy
import random
from torchvision import transforms
# import time
import torch.backends.cudnn as cudnn
import os, sys
from time import time, strftime

class MyResNet(nn.Module):
    def __init__(self, depth, num_classes, pretrained = True):
        super(MyResNet, self).__init__()
        if depth == 18:
            model = models.resnet18(pretrained)
        elif depth == 34:
            model = models.resnet34(pretrained)
        elif depth == 50:
            model = models.resnet50(pretrained)
        elif depth == 152:
            model = models.resnet152(pretrained)

        self.num_ftrs = model.fc.in_features
        # self.num_classes = num_classes

        self.shared = nn.Sequential(*list(model.children())[:-1])
        self.target = nn.Sequential(nn.Linear(self.num_ftrs, num_classes))

    def forward(self, x):
        # pdb.set_trace()

        x = self.shared(x)
        x = torch.squeeze(x)
        return self.target(x)

    def frozen_until(self, to_layer):
        print('Frozen shared part until %d-th layer, inclusive'%to_layer)

        # if to_layer = -1, frozen all
        child_counter = 0
        for child in self.shared.children():
            if child_counter <= to_layer:
                print("child ", child_counter, " was frozen")
                for param in child.parameters():
                    param.requires_grad = False
                # frozen deeper children? check
                # https://spandan-madan.github.io/A-Collection-of-important-tasks-in-pytorch/
            else:
                print("child ", child_counter, " was not frozen")
                for param in child.parameters():
                    param.requires_grad = True
            child_counter += 1





from sklearn.model_selection import train_test_split
# from utils import *
import torch
import torch.nn as nn
import argparse
import copy
import random
from torchvision import transforms
# import time
import torch.backends.cudnn as cudnn
import os, sys
from time import time, strftime

"""
Predict output of a new sample given model/models and fn/fns/dset
"""

def softmax_stable(Z):
    """
    Compute softmax values for each sets of scores in Z.
    each row of Z is a set of score.
    """
    e_Z = np.exp(Z - np.max(Z, axis = 1, keepdims = True))
    A = e_Z / e_Z.sum(axis = 1, keepdims = True)
    return A

def loader_len(dset):
    """
    return len of a DataLoader

    dset: a DataLoader that return (data, lbs, fns) at a batch
    """
    res = 0
    for _, _, fns in dset:
        res += len(fns)
    return res

def get_num_classes(model):
    """
    return number of output class of a pytorch classification model
    """
    # Step 1: unparalell_model
    model = unparallelize_model(model)
    return model.target[0].out_features


def singlemodel_score(model, dset_loader, num_tests = 1):
    """
    Use ONE pretrained model to predict score and probs each input in dset
    Make multiples predictions and accumulate results

    ----
    INPUT:
        model: a pretrained model
        dset: a MyDatasetSTFT variable
    OUTPUT:
        pred_outputs: np array -- prediction results, sum of all ouput before softmax
        pred_probs: np array -- prediction results, sum of all probability
        fns: list of filenames in loader order
    """
    num_files = loader_len(dset_loader)
#     num_classes = get_num_classes(model)
    num_classes = 2

    # preds = np.zeros((num_files, n_tests))
    total_scores = np.zeros((num_files, num_classes))
    total_probs = np.zeros((num_files, num_classes))
    fns = []
    torch.set_grad_enabled(False)
    model.eval()

    for test in range(num_tests):
        tot = 0
        print('test {}/{}'.format(test + 1, num_tests))
        start = 0
        for batch_idx, (inputs, labels, fns0) in enumerate(dset_loader):
            n = len(fns0)
            inputs = inputs.to(device)
            output = model(inputs)
            output = output.view((-1, num_classes))
            _, pred  = torch.max(output.data, 1)
            # preds[start:start + n, test] = pred.data.cpu().numpy()
            # pdb.set_trace()
            total_scores[start:start + n, :] += output.data.cpu().numpy()
            start += n
            tot += len(fns0)
            if test == 0: fns.extend(fns0)

        total_probs += softmax_stable(total_scores)
    return total_scores, total_probs, fns


def singlemodel_class(model, dset_loader, num_tests = 1):
    """
    predict label for dset_loader using one model
    This one is done after calling sm_score(model, dset_loader, num_tests)
    and get the torch.max(out.data, 1)
    """
    total_scores, total_probs, fns = singlemodel_score(model, dset_loader, num_tests)
    score_class = np.argmax(total_scores, axis = 1)
    prob_class = np.argmax(total_probs, axis = 1)
    return score_class, prob_class, fns

def multimodels_score(models, dset_loader, num_tests = 1):
    """
    accumulate results from multiple models
    output total scores and probabilities
    ----------
    INPUT:
        model: a list of pytorch classification models
    """
#     num_models = len(models)
    num_models =1
    num_files = loader_len(dset_loader)
    num_classes = get_num_classes(models)
    # preds = np.zeros((num_files, n_tests))
    total_scores = np.zeros((num_files, num_classes))
    total_probs = np.zeros((num_files, num_classes))
    for model in models:
        tmp_score, tmp_prob, fns = singlemodel_score(model, dset_loader, num_tests)
        total_scores += tmp_score
        total_probs += tmp_prob

    return total_scores, total_probs, fns


def multimodels_class(models, dset_loader, num_tests = 1):
    total_scores, total_probs, fns = multimodels_score(models, dset_loader, num_tests)
    score_class = np.argmax(total_scores, axis = 1)
    prob_class = np.argmax(total_probs, axis = 1)
    return score_class, prob_class, fns

def multimodels_multiloaders_score(models, dset_loaders, num_tests = 1):
    """
    accumulate results from multiple models, each model uses one dsetloader
    number of models == number dset_loaders
    output total scores and probabilities
    ----------
    INPUT:
        model: a list of pytorch classification models
    """
    assert len(models) == len(dset_loaders)
    num_models = len(models)
    num_files = loader_len(dset_loaders[0])
    num_classes = get_num_classes(models[0])
    # preds = np.zeros((num_files, n_tests))
    total_scores = np.zeros((num_files, num_classes))
    total_probs = np.zeros((num_files, num_classes))
    cnt = 0
    for model, dset_loader in zip(models, dset_loaders):
        cnt += 1
        print('######################')
        print('Model {}/{}'.format(cnt, len(models)))
        tmp_score, tmp_prob, fns = singlemodel_score(model, dset_loader, num_tests)
        # pdb.set_trace()
        total_scores += tmp_score
        total_probs += tmp_prob

    return total_scores, total_probs, fns

def multimodels_multiloaders_class(models, dset_loaders, num_tests = 1):
    total_scores, total_probs, fns =\
        multimodels_multiloaders_score(models, dset_loaders, num_tests)
    score_class = np.argmax(total_scores, axis = 1)
    prob_class = np.argmax(total_probs, axis = 1)
    return score_class, prob_class, fns
parser = argparse.ArgumentParser(description='PyTorch Digital Mammography Training')

parser.add_argument('--lr', default=1e-3, type=float, help='learning rate')
parser.add_argument('--net_type', default='resnet', type=str, help='model')
parser.add_argument('--depth', default=50, choices = [11, 16, 19, 18, 34, 50, 152, 161, 169, 121, 201], type=int, help='depth of model')
parser.add_argument('--weight_decay', default=5e-4, type=float, help='weight decay')
parser.add_argument('--finetune', '-f', action='store_true', help='Fine tune pretrained model')
parser.add_argument('--trainer', default='adam', type = str, help = 'optimizer')
parser.add_argument('--duration', default= 2.5, type = float, help='time duration for each file in second')
parser.add_argument('--n_tests', default=3, type = int, help='number of tests in valid set')
parser.add_argument('--gender', '-g', action='store_true', help='classify gender')
parser.add_argument('--accent', '-a', action='store_true', help='accent classifier')
parser.add_argument('--random_state', '-r', default = 2, type = int, help='random state in train_test_split')

parser.add_argument('--model_path', type=str, default = ' ')
parser.add_argument('--gamma', default = 0.5, type = float)
parser.add_argument('--batch_size', default=64, type=int)
parser.add_argument('--num_epochs', default=100, type=int,
                    help='Number of epochs in training')
parser.add_argument('--dropout_keep_prob', default=0.5, type=float)
parser.add_argument('--check_after', default=5,
                    type=int, help='check the network after check_after epoch')
parser.add_argument('--train_from', default=1,
                    choices=[0, 1],  # 0: from scratch, 1: from pretrained 1 (need model_path)
                    type=int,
                    help="training from beginning (1) or from the most recent ckpt (0)")

parser.add_argument('--frozen_until', '-fu', type=int, default = -1,
                    help="freeze until --frozen_util block")
parser.add_argument('--val_ratio', default=0.1, type=float,
        help = "number of training samples per class")

########################################################################################33

args, unknown = parser.parse_known_args()



# from __future__ import division
# from split import my_split
# from sklearn.model_selection import train_test_split
# from myloss import SmoothLabel
from sklearn.metrics import accuracy_score
# import config as cf
# from nets import MyResNet
# import utils
# from data import build_dataloaders
import torch

import torch.nn as nn
import argparse
import copy
import random
from torchvision import transforms
# import time
import torch.backends.cudnn as cudnn
import os, sys
from time import time, strftime
import pdb
from scipy.io import wavfile
from sklearn.metrics import confusion_matrix
# from mytrain_test_split import mytrain_test_split_voice
import warnings
warnings.filterwarnings("ignore", message="numpy.dtype size changed")
warnings.filterwarnings("ignore", message="numpy.ufunc size changed")
# from predicts import singlemodel_class
# from nets import load_model, parallelize_model


print('======================================================')
print('Data preparation')

def exp_lr_scheduler(args, optimizer, epoch):
    # after epoch 100, not more learning rate decay
    init_lr = args.lr
    lr_decay_epoch = 4 # decay lr after each 10 epoch
    weight_decay = args.weight_decay
    lr = init_lr * (0.6 ** (min(epoch, 200) // lr_decay_epoch))

    for param_group in optimizer.param_groups:
        param_group['lr'] = lr
        param_group['weight_decay'] = weight_decay

    return optimizer, lr

saved_models = './saved_model/'
if not os.path.isdir(saved_models): os.mkdir(saved_models)
saved_model_fn = saved_models + args.net_type + '_' +\
    str(args.depth) + '_' +  strftime('%m%d_%H%M') + '_r' + str(args.random_state)
print('model will be saved to {}'.format(saved_model_fn))
print('********************************************************')
def load_model(path_model):
    old_model = path_model
    if args.train_from == 1 and os.path.isfile(old_model):
        print("| Load pretrained at  %s..." % old_model)
        checkpoint = torch.load(old_model, map_location=torch.device('cpu'))
        tmp = checkpoint['model']
        model = unparallelize_model(tmp)
        try:
            top1acc = checkpoint['acc']
            print('previous acc\t%.4f'% top1acc)
        except KeyError:
            pass
        print('=============================================')
    else:
        model = MyResNet(args.depth, 2)

    return model










