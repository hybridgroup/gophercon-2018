# Intel Movidius Myriad 2 Neural Compute Stick (NCS)

The Intel Movidius Myriad 2 Neural Compute Stick (https://www.movidius.com/myriad2) is a Video Processing Unit (VPU) that lets you perform low power execution of deep neural networks, in the form of a USB device.

## Installation

You must first install the NCSDK before you can use the Intel Movidius Myriad 2 Neural Compute Stick. The official SDK only supports Linux. However, the fork created by [@milosgajdos83](https://github.com/milosgajdos83) has some initial support for macOS. Sorry, no Windows yet.

### macOS

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
    cd api/src && sudo make basicinstall

## Code

### step01/main.go

First, let's just verify communication with the NCS.

### step02/main.go

Now we will load a Caffe deep neural network graph on to the NCS to process an image file.

### step03/main.go

We can use the input for an attached webcam to perform image classification using the Caffe deep neural network graph we used in the previous step.
