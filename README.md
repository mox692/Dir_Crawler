## ABOUT
・Dir_Crawler is directory croler.
・You can find files by keywords that is in file or filename.

## install
Access the [release page](https://github.com/mox692/ChromeExtention_SettingTimer/releases) or `git clone`.  

## Usage
When you use dir_crawler, I recommend you to run binary through shellscript which is in ***`_toos`*** directory,  
so that you can use `--jump` command.

## subcommands

![dirclole_flow](https://user-images.githubusercontent.com/55653825/98936227-0653b200-2528-11eb-82af-bf1058a283e9.png)

**There are 2 sub commands.**
### list
This is the basic command.  
`list` lists the files that have the name specified by kw.  
```
# you can jump to the dir where serching file is in.
$crawl list --kw="filename"
```

### jump
`jump` jumps to the dir where serching file is in.  
```
##you can jump to the dir where serching file is in.
$crawl jump --kw="filename"
```

### get
`get` gets the serching file by copying. Copied file will be set in current dir.
```
##you can find the file and copy it to current dir. 
$crawl jump --kw="filename"
```
# flags

## --kw(keyword)
In anytime, `kw` is the required option.Crowler finds the file with the name specified by kw.  
Crawler also find the file by filecontents that is corresponding to kw.

