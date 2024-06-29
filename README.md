# Viola Ransomware

Viola Ransomware is a software for simulating ransomware that uses AES-256-GCM encryption & created using the Golang language.

## Features

- Support Windows & Linux
- Ransom Message
- Encrypt & Decrypt Files
- Etc 

## Install
Install Go Language First. Click <a href="https://go.dev/doc/install">here</a> for download<br>
1. Clone/Download this repository
```
git clone https://github.com/alf4ridzi/viola-ransomware
```
2. Build Executable File
```
go build .
```
3 Run the executable file
```
viola-ransomware.exe
```

## Usage Tips
- This ransomware will encrypt certain file extensions that have been set in the config/config.go section.
- Go to config/config.go to set the config before running the ransomware<br>
- To decrypt files, select option 2 (decrypt files) & enter the key in the key.txt file.

## Videos
<b>Encrypt</b><br>
![encrypt](https://github.com/alf4ridzi/viola-ransomware/assets/58920998/f29a1f7f-9a43-47fb-8bda-3805ded646c9)<br>
<b>Decrypt</b><br>
![decrypt](https://github.com/alf4ridzi/viola-ransomware/assets/58920998/e4f21377-6717-4541-97d9-3e9237793151)

## License
Release in 2024 under <a href="https://opensource.org/license/mit">MIT License</a>
