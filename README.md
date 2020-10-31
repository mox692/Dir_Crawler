# ABOUT
・Dir_Crawler is directory croler.
・You can find files by keywords that is in file or filename.

# get binary
Access the [release page](https://github.com/mox692/ChromeExtention_SettingTimer/releases) and download the project folder. 

# crawl subcommands
**There are 3 sub commands.**
## list
This is the basic command.  
`list` lists the files that have the name specified by kw.  
```
# you can jump to the dir where serching file is in.
$crawl list --kw="filename"
```

## jump
`jump` jumps to the dir where serching file is in.  
```
# you can jump to the dir where serching file is in.
$crawl jump --kw="filename"
```

## get
`get` gets the serching file by copying. Copied file will be set in current dir.
```
# you can find the file and copy it to current dir. 
$crawl jump --kw="filename"
```
# crawl flags

## --kw(keyword)
In anytime, `kw` is the required option.Crowler finds the file with the name specified by kw.  
Crawler also find the file by filecontents that is corresponding to kw.
