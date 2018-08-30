# Installing on Kali 2018.2(unsupported platform)

From the terminal:

   * Create your work directories ```mkdir -p ~/workspace/ncsdk```
   * Download 2.05.00.02 from https://github.com/movidius/ncsdk to ~/workspace
   * Extract the archive from ncsdk ```cd ~/workspace/ncsdk; tar zxvf ../ncsdk-2.05.00.02.tar.gz --strip-components 1```
   * Install the dependencies required to compile ```sudo apt-get install -y libusb-1.0-0-dev libprotobuf-dev libleveldb-dev libsnappy-dev libopencv-dev libhdf5-serial-dev protobuf-compiler libatlas-base-dev git automake byacc lsb-release cmake libgflags-dev libgoogle-glog-dev liblmdb-dev swig3.0 graphviz libxslt-dev libxml2-dev gfortran python3-dev python-pip python3-pip python3-setuptools python3-markdown python3-pillow python3-yaml python3-pygraphviz python3-h5py python3-nose python3-lxml python3-matplotlib python3-numpy python3-dateutil python3-skimage python3-scipy python3-six python3-networkx```
   * Install the python3 dependencies required for some of the examples ```pip3 install {graphviz,networkx,,scikit-image,google,protobuf,tensorflow}```
   * Edit the install script ```vi ~/workspace/ncsdk/install.sh```, comment lne 43 (exit 1 after the message 'Your current combination of Linux distribution and distribution version is not officially supported! Error on line $LINENO.  Will exit'). This only works because Kali is based off of upstream Debian.
   * Run ```cd ~/workspace/ncsdk; install.sh```, it should complete without errors

At this point the sdk and tools should be installed and working. Close your terminal and reopen it for all the exports to reload. 

Plug in the Movidius chip then download the test apps from https://github.com/movidius/ncappzoo to validate things are working fine.
