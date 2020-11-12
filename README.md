## ABOUT
・Dir_Crawler is directory croler.
・You can find files by keywords that is in file or filename.
・You can also jump to directory where finding file is in.

## install
Access the [release page](https://github.com/mox692/ChromeExtention_SettingTimer/releases) or `git clone`.  

## Usage
When you use dir_crawler, I recommend you to run binary through shellscript which is in ***`_toos`*** directory,  
so that you can use `--jump` command.  
![dirclole_flow](https://user-images.githubusercontent.com/55653825/98936227-0653b200-2528-11eb-82af-bf1058a283e9.png)  

## subcommands
**There are 2 sub commands.**

### list
`list` lists files that match the keyword.  
```
##you can find the files that match keyword. 
$crawl --list="keyword"
```

### get
`jump` jumps to the dir where serching file is in.  
```
##you can jump to the dir where serching file is in.  
$crawl --jump="keyword"
```
# flags

## --kw(keyword)
In anytime, `kw` is the required option.Crowler finds the file with the name specified by kw.  
Crawler also find the file by filecontents that is corresponding to kw.

