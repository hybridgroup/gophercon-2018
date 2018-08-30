# Intel Movidius Myriad 2 Neural Compute Stick (NCS)

The Intel Movidius Myriad 2 Neural Compute Stick (https://developer.movidius.com/) is a Video Processing Unit (VPU) that lets you perform low power execution of deep neural networks, in the form of a USB device.

## Installation

You must first install the NCSDK before you can use the Intel Movidius Myriad 2 Neural Compute Stick. The official SDK only supports Linux. However, the fork created by [@milosgajdos83](https://github.com/milosgajdos83) has some initial support for macOS. Sorry, no Windows yet.

### macOS

The macOS support for the NCSDK currently is only for the API, not the graph compiler or other tools. To install, run the following commands:

    brew install coreutils opencv libusb pkg-config wget
    git clone https://github.com/milosgajdos83/ncsdk.git
    cd ncsdk
    git checkout macos-V1
    cd api/src && sudo make basicinstall

### Linux

You must have OpenCV and GoCV installed in order to use the Movidius SDK with Go:

https://gocv.io/getting-started/linux/

Once they are installed, you can run the following commands:

    git clone https://github.com/milosgajdos83/ncsdk.git
    cd ncsdk
    make install

## Precompiled models

There is a Dropbox folder that you can download precompiled graph files for each of the models that are included in the NCSDK examples. Not all of these models have corresponding examples in the `go-ncs` package yet. You can find the graph files here:

https://www.dropbox.com/sh/gaxc0sb1c1n54q8/AAAz27hbwos5WtZi_j5j9qSza?dl=0

## Code

### step1/main.go

First, let's just verify communication with the NCS.

	go get github.com/hybridgroup/go-ncs
	go run main.go 0 

### step2/main.go

Now we will load a Caffe deep neural network graph on to the NCS to process an image file.

	go get -u gocv.io/x/gocv

Use graph files from the dropbox folder specified in the parent README.
Sample image can be obtained from 'ncsdk' you cloned/downloaded in the previous step.
The description file is part of 'ncsdk' as well. However, you need to build/generate it

    cd </path/to/ncsdk>/examples/data/ilsvrc12
    make

    go run main.go 0 </path/to/dropbox>/ncs-models/caffe/googlenet/graph </path/to/ncsdk>/examples/data/images/cat.jpg </path/to/ncsdk>/examples/data/ilsvrc12/synset_words.txt
    
### step3/main.go

We can use the input from an attached webcam to perform image classification using the Caffe deep neural network graph we used in the previous step.

### step4/main.go

Now, let's use the input from an attached webcam to perform image classification using the Tensorflow Inception v3 deep neural network model.
