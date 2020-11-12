## ABOUT
・Dir_Crawler help you to find files when you lost the path of it.   
・You can find files by keyword which is in file or filename.  
・You can also jump to directory where finding file is in.  

## Install
Access the [release page](https://github.com/mox692/ChromeExtention_SettingTimer/releases) or `git clone`.  

## Usage
When you use dir_crawler, I recommend you to run binary through shellscript which is in ***`_toos`*** directory,  
so that you can use `--jump` command.  
![dirclole_flow](https://user-images.githubusercontent.com/55653825/98936227-0653b200-2528-11eb-82af-bf1058a283e9.png)  

## Setup

### set shellscript   
*`./_tools/cw`*
```
...
### !バイナリ(dirWalk)のpathに合わせて変更してください 
result=$(path/to/binary --$action=$kw)
status=$? 
...
```

### Create an alias for `cw` to run easily.  
*`terminal`*
```
$ alias (command_name)='path/to/cw'
```

## Commands
There are 2 sub commands.  

### list
`list` lists files that match the keyword.  
```
## shellscript (./_tools/cw)
$cw list keyword

## go binary. 
$crawl --list="keyword"
```

### jump
`jump` jumps to the dir where serching file is in.  conta
```
## shellscript (./_tools/cw)
$cw list keyword
```

## contact
If you notice something wrong, please send a issue:)  

## License  
MIT

