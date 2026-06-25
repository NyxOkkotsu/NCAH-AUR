# 🐾 NCAH : Nyx Cutest AUR Helper 🐾

<pre>
  _   _  ____    _    _   _ 
 | \ | |/ ___|  / \  | | | |
 |  \| | |     / _ \ | |_| |
 | |\  | |___ / ___ \|  _  |
 |_| \_|\____/_/   \_\_| |_|
                            
 🐾 NCAH : Nyx Cutest AUR Helper 🐾
</pre>
</p>

---

## ✨ What is this meww?
Hwaaa~ Welcome to **NCAH**! This AUR Helper is Focused on security and PKG transparency it can scan the PKGBUILD for u nyaww~! (✿•ᴗ•)

Nyx got super tired of boring helpers that blindly install scary, unpredictable stuff into your beautiful system, so Nyx built this from scratch using pure Go! It keeps your Arch Linux system 100% safe from mean, naughty scripts while staying absolutely adorable and fluffy!

---

## 📺 Video Preview
Look how cute and incredibly smart Nyx is when protecting your terminal from bad packages! Check out the quick show nyaa~ 👇

[https://github.com/user-attachments/assets/00000000-0000-0000-0000-000000000000](https://github.com/user-attachments/assets/00000000-0000-0000-0000-000000000000)


---

## 🛡️ Super Smart Security Scanning!
Before installing anything, Nyx will perform a magical **Security X-Ray Scan** on the target `PKGBUILD` and sniff out bad, suspicious things like:
* 🙀 **Dangerous Execution Pipelines:** `curl | bash` or `wget | sh` hidden in the script!
* 😭 **System Destructors:** The extremely scary `rm -rf /` ghost!
* 😤 **Sneaky Sudo:** Explicit `sudo` usage declared inside the PKGBUILD routines.
* 🌐 **Unencrypted Sources:** Non-HTTPS URLs (`http://`) that eavesdroppers can spy on!
* 🤔 **Secret Obfuscations:** Strange `base64 -d` decoding tricks or sneaky `eval` usage.
* 💢 **Missing Integrity Checks:** Packages trying to run away with `sha256sums=('SKIP')`!

Then, Nyx will instantly classify them into **SAFE** ✨, **WARNING** ⚠️, or **HIGH RISK** 🔥🙀 levels to protect your home!

---

## 🚀 How to Build & Install

Make sure you already have go and git installed on your computer, okay? (๑•̀ㅂ•́)و✧
Bash

```text

# 1. Clone your lovely repository
git clone https://github.com/YOUR_GITHUB_USERNAME/ncah.git
cd ncah

# 2. Tidy up the Go modules
go mod tidy

# 3. Build the magical executable binary!
go build -o ncah cmd/ncah/main.go

```

If you want to invoke Nyx from anywhere across your system without typing ./, simply move the binary to your system PATH:
Bash

sudo mv ncah /usr/local/bin/ncah

## 🛠️ How to Use nyaaa~~!

🔍 1. Search for Cute Packages

```text
ncah -Ss pfetch
```

🫣 2. Peek at Package Info

```text
ncah -Si uwufetch
```

🛡️ 3. Securely Install a Package

```text
ncah -S pipes.sh
```

## 📁 Project Structure
Nyx keeps her room and her project workspace super tidy and organized meww!
```text
ncah/
├── cmd/
│   └── ncah/
│       └── main.go            # The adorable brain of the helper!
├── internal/
│   ├── aur/                   # Fetching, searching, and peeking from the AUR API
│   ├── build/                 # Compiling and assembling with love (makepkg)
│   ├── security/              # The X-Ray scanner mechanics & clever safety rules 🛡️
│   ├── ui/                    # Beautiful ASCII banners & colorful prompt dialogs
│   └── utils/                 # Secret utility helpers
├── pkg/
│   └── logger/                # Logging buddies
├── go.mod                     # Go dependency module map
└── README.md                  # You are reading this right now nyaa~


