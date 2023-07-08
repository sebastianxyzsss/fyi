# FYI

this cli tool takes any markdown and serves it as an html document so it can be opened in any browser. However, any quoted string or anything inside a code block will be converted as copyable content, meaning that it can be clicked and it will be copied to the clipboard.

## Instructions

to start serving the markdown file:
```
fyi <markdown file>
```
then go to any browser and open the url http://localhost:8888/

# Installation

For now run makefile, the executable gets create in the build folder. I will create a debian package in the future.
