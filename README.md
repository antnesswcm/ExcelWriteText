# ExcelWriteText

## 项目介绍

ExcelWriteText是一个命令行工具,用于将Excel表格中的文本按行写入对应文件。例如,如果a列为文件名,b列为文件内容,该工具可以将b列的内容写入a列对应的文件中。

## 使用方法

```
src.exe [选项] [files]  
```  

## 选项

- -b    遇见空单元停止  
- -c string                  
    以逗号分隔的列地址列表如:"-c a,b"                
- -e string                  
    指定文件后缀名 (default "txt")                    
- -f string                  
    以逗号分隔的文件列表,不可与[files]同时使用                  
- -h    显示帮助信息             
- -help                       
    显示帮助信息           
- -o string                  
    指定输出文件夹 (default "./output")          
- -s string                  
    以逗号分隔的列地址范围如:"-s a-b,c-d"

## 示例

将excel文件中a列的文件名和b列的内容分别写入对应文件:

```
src.exe -f files.xlsx 
```  

输出文件保存到output文件夹:

```
src.exe -f files.xlsx -o d:\output\
```  

只处理a-b列和c-d列:

```
src.exe -f files.xlsx -s a-b,c-d
```  

遇到空单元格停止:

``` 
src.exe -f files.xlsx -b
```

指定输出文件为.txt扩展名:

```
src.exe -f files.xlsx -e txt
```

## 联系方式

有任何问题可以通过issues联系我。