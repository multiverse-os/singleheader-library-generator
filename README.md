[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse OS: Single-file header only library geneerator
**URL** [multiverse-os.org](https://multiverse-os.org)

This Go single-file header only library generator allows developers to
establish a YAML configuration for a C or C++ library. Either for C or C++
library developers to offer their library in different forms to make it easier
to use; or for developers wishing to include a C or C++ library in their project
but want it packaged in a specific way, and importantly, be able to regenerate
this single-file header only version whenever changes are made upstream.  

Several ongoing experiments for Multiverse OS require C and C++ libraires and
this is becomming our preferred way to package these libraries in our Go
projects. The goal of this project is to be more broadly useful for a wider
variety of use-cases than just our own specific use case because in our search
for a similar tool, we were not able to find many options, and they did not
provide us with much flexibility. 

#### Usage 
The structure of the YAML file is not yet defined, because this project is under
active development. 

Multiverse OS developers encourage compiling projects over using pre-compiled 
binaries, and so we will include easy set of instructions for compiling the
binary and generating YAML configurations for a few libraries to illustrate the
entire process which should be simple and ideally a few steps. 

