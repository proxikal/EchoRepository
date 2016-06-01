package gto

import (
	"io"
	"os"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"time"
	"fmt"
	"runtime"
	"bufio"
	"math/rand"
	"path/filepath"
	"os/exec"
	"bytes"
	"strings"
	"errors"
	"strconv"
	"net/url"
	//"log"
	"encoding/base64"
	"encoding/binary"


	"github.com/TrevorSStone/goriot"
	"github.com/bwmarrin/discordgo"
	"github.com/ChimeraCoder/anaconda"
)



	type system struct {
		Debug		bool
		Speak		bool
	}

	type error interface {
    	Error() string
	}

	type errorString struct {
    	s string
	}


	type pushString struct {
		s string
	}


	type Set struct {
		Give 	string
		Take 	string
		Color 	string
	}


	var Command Set // set's the command names for Gotools Discord System.
	var Discord system // Global variable to enable Debug or Speak. (Not yet implemented)
	var buffer = make([][]byte, 0) // variable for playsound() function for playing .dca music over discord.
	var GiveErrorResponse string // Error response for the command "Give"
	var TakeErrorResponse string // Error Response for the command "Take"
	var ColorErrorResponse string // Error Response for the command "Rolecolor"
	var User string // variable to return the username. (for commands give, take)
	var Role string // variable to return the role. (for commands give, take, rolecolor)
	var Master string // variable to set the master role. (whoever has this role has control over Gotools)
	var Color string // variable to return the color (for commands rolecolor)
	var err error



	// let's set variables for the Autorole system
	var AutoRole bool
	var AutoRoleName string
	var AutoRoleBots bool



	// let's set the variables for Greet & Leave Message
	var Greet bool
	var GreetMessage string
	var Leave bool
	var LeaveMessage string
	var PmGreet bool
	var PmGreetMessage string


func (e *errorString) Error() string {
    return e.s
}


func New(text string) error {
    return &errorString{text}
}





// We're using anaconda for Twitter Stuff!
func ConsumerKey(key string) {
	anaconda.SetConsumerKey(key)
}


func ConsumerSecret(key string) {
	anaconda.SetConsumerSecret(key)
}


func TwitterApi(token string, secret string) *anaconda.TwitterApi {
	return anaconda.NewTwitterApi(token, secret)
}




func TwitterSearch(keyword string, opt int, api *anaconda.TwitterApi) []string {
	var tweets []string
	v := url.Values{}
	opti, err := String(opt)
	if err != nil {
		Print("Gotools Error: function TwitterSearch() you need to type an amount TwitterSearch(keyword, amount, api)")
		return nil
	}
	v.Set("count", opti)
	searchResult, err := api.GetSearch(keyword, v)
	for _ , tweet := range searchResult.Statuses {
    	// fmt.Print(tweet.Text)
    	tweets = append(tweets, tweet.Text)
	}
	return  tweets
}





// League of Legends using: goriot
// personalkey = os.Getenv("THEIRKEYHERE")

func RiotKey(key string) {
	goriot.SetAPIKey(key)
}


func RiotByName(region string, user string) (map[string]goriot.Summoner, error) {
	fruit, err := goriot.SummonerByName(region, user)
	return fruit, err
}



func GetRiotID(region string, user string) int {
  data := 0
  RiotKey("bf4d4d68-9b08-4591-855d-6763a2ec5c62")
  sum, err := RiotByName(region, user)
  if err != nil {
//    t.Error(err.Error())
 //   fmt.Println(err)
    return 0
  }
//  fmt.Println(sum)
  for _, v := range sum {
    if v.Name == user {
      userid := fmt.Sprintf("%d", v.ID)
    //  icon := fmt.Sprintf("%d", v.ProfileIconID)
     //  level := strconv.Itoa(v.SummonerLevel)
    //  fmt.Println("User ID:" + userid)
      data, err = strconv.Atoi(userid)
      if err != nil {
        return 0
      }
    }
  }
//  theid := strconv.Itoa(sum.ID)
//  fmt.Println("User:"+sum.Name + "\nID:"+theid)
  return data
}



func RiotSmallLimit(requests int, limit time.Duration) {
  // goriot.SetSmallRateLimit(10, 10*time.Second)
	goriot.SetSmallRateLimit(requests, limit)
}

func RiotLargeLimit(requests int, limit time.Duration) {
	// goriot.SetLongRateLimit(500, 10*time.Minute)
  goriot.SetLongRateLimit(requests, limit)
}



func ReadDir(path string, ext string) ([]string, error) {
	path = strings.TrimSuffix(path, "/")
	data, err := filepath.Glob(path+"/*"+ext)
	if err != nil {
		return nil, err
	}
	return data, err
}



func CleanFileName(path string, ext string, opts []string) error {
	path = strings.TrimSuffix(path, "/")
	file, err := filepath.Glob(path+"/*"+ext)
	if err != nil {
		return err
	}

	for _, v1 := range file {
		for _, v := range opts {
			newv := strings.Replace(v1, v, "", -1)
			err = os.Rename(v1, newv)
			if err != nil {
				return errors.New("rename error on file: " + v1)
			}	
		}
	}
	return err
}






func CleanPath(path string) string {
	newname := Replace(path, ",", "", -1)
	newname = Replace(newname, " ", " ", -1)
	newname = Replace(newname, "\"", "", -1)
	newname = Replace(newname, ":", "", -1)
	newname = Replace(newname, "`", "", -1)
	newname = Replace(newname, "^", "", -1)
	newname = Replace(newname, "*", "", -1)
	newname = Replace(newname, "!", "", -1)
	newname = Replace(newname, "/", "", -1)
	newname = Replace(newname, "\\", "", -1)
	newname = Replace(newname, "?", "", -1)
	newname = Replace(newname, "<", "", -1)
	newname = Replace(newname, ">", "", -1)
	newname = Replace(newname, "|", "", -1)
	newname = Replace(newname, " _ ", " ", -1)
	newname = Replace(newname, "  ", " ", -1)
	return newname
}






func Replace(str string, from string, to string, limit int) string {
	return strings.Replace(str, from, to, limit)
}




func Contains(haystack string, needle string) bool {
	return strings.Contains(haystack, needle)
}



func ToLower(str string) string {
	return strings.ToLower(str)
}


func ToUpper(str string) string {
	return strings.ToUpper(str)
}


func TrimPrefix(str string, prefix string) string {
	return strings.TrimPrefix(str, prefix)
}


func TrimSuffix(str string, suffix string) string {
	return strings.TrimSuffix(str, suffix)
}

func HasPrefix(str string, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

func HasSuffix(str string, suffix string) bool {
	return strings.HasPrefix(str, suffix)
}


func String(str int) (string, error) {
	// return &pushString{str}
	i := strconv.Itoa(str)
	return i, nil
}


func Integer(str string) (int, error) {
	// return &pushString{str}
	i , err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return i, nil
}



func Print(str interface{}) {
	fmt.Println(str)
}





func OpenURL(url string) {
	switch runtime.GOOS {
		case "linux":
   			err = exec.Command("xdg-open", url).Start()
		case "windows", "darwin":
   			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		default:
    	err = fmt.Errorf("can't open url. unsupported platform.")
	}

}




func DownloadFile(output string, url string) bool {
	chk := true
	os.Remove(output)
	out, err := os.Create(output)
	if err != nil {
		chk = false
	// fmt.Println(err)
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		chk = false
	// fmt.Println(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		chk = false
	// fmt.Println(err)
	}
	return chk
}






func GetJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()
    return json.NewDecoder(r.Body).Decode(target)
}





func Unmarshal(data []byte, target interface{}) error {
	return json.Unmarshal(data, &target)
}





func Marshal(target interface{}) ([]byte, error) {
	jso, err := json.MarshalIndent(target, "", "   ")
	if err != nil {
		return nil, err
	}
return jso, nil
}






func ReadFile(path string) ([]byte, error) {
	tb, err := ioutil.ReadFile(path)
	return tb, err
}



func WriteFile(path string, b []byte, perm os.FileMode) {
	ioutil.WriteFile(path, b, perm)
}



func Random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}





func Split(str string, op string) []string {
	return strings.Split(str, op)
}





func ReadLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}





func CountLines(path string) int {
  counter := 0

  file, err := os.Open(path)
  if err != nil {

  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    counter++
  }
  return counter
}





func copyFileContents(src, dst string) (err error) {
    in, err := os.Open(src)
    if err != nil {
        return
    }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil {
        return
    }
    defer func() {
        cerr := out.Close()
        if err == nil {
            err = cerr
        }
    }()
    if _, err = io.Copy(out, in); err != nil {
        return
    }
    err = out.Sync()
    return
}



func CopyFile(src, dst string) (err error) {
    sfi, err := os.Stat(src)
    if err != nil {
        return
    }
    if !sfi.Mode().IsRegular() {
        // cannot copy non-regular files (e.g., directories,
        // symlinks, devices, etc.)
        return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
    }
    dfi, err := os.Stat(dst)
    if err != nil {
        if !os.IsNotExist(err) {
            return
        }
    } else {
        if !(dfi.Mode().IsRegular()) {
            return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
        }
        if os.SameFile(sfi, dfi) {
            return
        }
    }
    if err = os.Link(src, dst); err == nil {
        return
    }
    err = copyFileContents(src, dst)
    return
}




func ifExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	} else {
		return true
	}
	return false
}





















// ################ DiscordGo Manipulation Below #################










func AutoRoleListen(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	if AutoRoleName == "" {
	//	Print("Gotools Error: function AutoRoleListen() You need to define what role you want me to add people to in your config.json file.")
		return
	}

	if AutoRole == false {
	//	Print("Gotools Error: function AutoRoleListen() You need to enable the Autorole system")
		return
	}

	if AutoRoleSystem(s, m) == true {
		// it works
	} else {
		Print("Gotools Error: function AutoRoleListen() An error occured while trying to autorole someone.")
	}
}











func CheckMaster(s *discordgo.Session, m *discordgo.MessageCreate, guildID string, author string) bool {
	if Master == "" {
		Print("Gotools Warning: Master variable not set.\nyou need to set gto.Master to your Bot Commander role.")
		return false
	}

	// Check if the user is Master.
	if MemberHasRole(s, guildID, author, Master) == true {
		return true
	}

	return false
}







func SendFile(s *discordgo.Session, m *discordgo.MessageCreate, path string) error {
	mk, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
    tc := bytes.NewReader(mk)
    s.ChannelFileSend(m.ChannelID, path, tc)
    return nil
}



func PostRandomImage(cat string, s *discordgo.Session, m *discordgo.MessageCreate, path string) error {
	img, _ := ioutil.ReadDir(path)
	cnt := len(img)
	rand := Random(1, cnt)
	path = TrimSuffix(path, "/")
	mk, err := ReadFile(path+"/"+img[rand].Name())
	if err == nil {
        tc := bytes.NewReader(mk)
        s.ChannelFileSend(m.ChannelID, cat+".png", tc)
	} else {
		return err
	}
	return nil
}





func Connect(token string) (*discordgo.Session, error) {
    // Login to discord. You can use a token or email, password arguments.
	return discordgo.New(token)
}





func ChangeAvatar(s *discordgo.Session, path string) bool {
	// grab user information
    img, err := ioutil.ReadFile(path)
    if err != nil {
    	fmt.Println(err)
    	return false
    }

    base64 := base64.StdEncoding.EncodeToString(img)
    avatar := fmt.Sprintf("data:image/png;base64,%s", string(base64))
    _, err = s.UserUpdate("", "", s.State.User.Username, avatar, "")
    if err != nil {
    	fmt.Println(err)
    }
    return false
}



func ChangeName(s *discordgo.Session, name string) bool {
    _, err = s.UserUpdate("", "", name, s.State.User.Avatar, "")
    if err != nil {
    	fmt.Println(err)
    	return false
    } else {
    	return true
    }
    return false
}




func Listen() {
	// Simple way to keep program running until any key press.
	var input string
	fmt.Scanln(&input)
	return
}






func SendMessage(s *discordgo.Session, channelID string, text string) {
	s.ChannelMessageSend(channelID, text)
}






func GetRoleName(s *discordgo.Session, guildID string, role string) string {
  var re string
  roles, err := s.State.Guild(guildID)
  if err == nil {
    for _, v := range roles.Roles {
      if v.ID == role {
        re = v.Name
      }
    }
  }
  return re
}






func GetRoleID(s *discordgo.Session, guildID string, role string) string {
  var re string
  roles, err := s.State.Guild(guildID)
  if err == nil {
    for _, v := range roles.Roles {
      if v.Name == role {
        re = v.ID
      }
    }
  }
  return re
}












// loadSound attempts to load an encoded sound file from disk.
func loadSound(path string) error {
	buffer = nil
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return err
	}

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}
}





// playSound plays the current buffer to the provided channel.
func playSound(s *discordgo.Session, guildID, channelID string) (err error) {
	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	// Sleep for a specified amount of time before playing the sound
	time.Sleep(250 * time.Millisecond)

	// Start speaking.
	vc.Speaking(true)

	// Send the buffer data.
	for _, buff := range buffer {
		vc.OpusSend <- buff
	}

	// Stop speaking
	vc.Speaking(false)

	// Sleep for a specificed amount of time before ending.
	time.Sleep(250 * time.Millisecond)

	// Disconnect from the provided voice channel.
	vc.Disconnect()
	buffer = nil
	vc.OpusSend = nil

	return nil
}









func MemberHasRole(s *discordgo.Session, GuildID string, AuthorID string, role string) bool {
  var opt bool
  opt = false
  therole := GetRoleID(s, GuildID, role)
z, err := s.State.Member(GuildID, AuthorID) 
if err != nil {
z, err = s.GuildMember(GuildID, AuthorID)
}
// }

if err == nil {
  var l []string
  l = z.Roles
  for r := range z.Roles {
    if therole == l[r] {
      opt = true
    }
  }
} // end of err == nil
  return opt
}








func ServerID(s *discordgo.Session, m *discordgo.MessageCreate) (string, error) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// Could not find channel.
		return "", err
	}
	return c.GuildID, nil
}










func PrintServerInfo(s *discordgo.Session) ([]byte, error) {
// 	var Info map[string]int
	servercount := 0
	rolecount := 0
	channelcount := 0
	membercount := 0
	g, err := s.UserGuilds()
	if err == nil {
		for _, v := range g {
			servercount++
			r, err := s.State.Guild(v.ID)
			if err == nil {
				for _ = range r.Roles {
					rolecount++
				}

				for _ = range r.Channels {
					channelcount++
				}

				for _ = range r.Members {
					membercount++
				}
			}
		}
}


	jso, err := json.Marshal(struct{
		Servers		int
		Roles 		int
		Channels 	int
		Members 	int
	} {
		Servers: servercount,
		Roles: rolecount,
		Channels: channelcount,
		Members: membercount,
	})
	if err != nil {
		return nil, err
	}
return jso, nil
}












func TakeRole(s *discordgo.Session, m *discordgo.MessageCreate, prefix string, cmd string) bool {
	if cmd == "" {
		Print("Gotools Error: function TakeRole() You need to declare a Command name\nUse gto.SetName.Take to set a name.")
		return false
	}
	if prefix == "" {
		Print("Gotools Error: function TakeRole() You need to declare a prefix.\nMake sure you have the prefix variable set in main() function.")
		return false
	}

	str := Replace(m.Content, prefix + cmd + " ", "", -1) // replaces the prefix and command name from the string.

	if Contains(str, " ") == false { // Look to see if they have the user id AND the role name.
		SendMessage(s, m.ChannelID, TakeErrorResponse)
		return false
	}



	if Contains(str, "<@") == false { // Let's make sure an ID Exists to prevent any errors.
		SendMessage(s, m.ChannelID, TakeErrorResponse)
		return false
	}




	// Now we need to clean this up a bit.
	data := Split(str, "<@") // first out of three steps to cleaning the id.
	a := Split(data[1], ">") // second step notice the data[1]
	therole := Replace(m.Content, prefix + "take <@"+a[0]+"> ", "", -1) // clean the original string to get the Role Name.
	theuser := Replace(a[0], "!", "", -1) // nickname support. If they have a nickname it adds an !
	User = theuser
	Role = therole

	var roleID string
	guildID, err := ServerID(s, m)
	mroles, err := s.State.Guild(guildID)
	ck := 0
	if err == nil {
 		for _, v := range mroles.Roles {
    		if v.Name == therole {
    			ck++
    			roleID = v.ID
    		}
  		}
  	} else {
  		Print("GoTools function TakeRole() Error:")
  		Print(err)
  		return false
  	}


  	if ck == 0 {
  		return false
  	}


	x, err := s.State.Member(guildID, theuser) 
	if err != nil {
		Print("Gotools function TakeRole() Error:")
		Print(err)
	} else {
		var ms []string
		ms = x.Roles
		for mr := range x.Roles {
			t := ms[mr]
			if t == roleID {
    			x.Roles = append(x.Roles[:mr], x.Roles[mr+1:]...)
    			s.GuildMemberEdit(guildID, theuser, x.Roles)
    			return true // everything worked. and the role was removed.
			}
		}
	}
	return false
}






func RoleColor(s *discordgo.Session, m *discordgo.MessageCreate, prefix string, cmd string) bool {
	var roleID string
	var hoist bool
	var perms int
	if cmd == "" {
		Print("Gotools Error: function TakeRole() You need to declare a Command name\nUse gto.SetName.Take to set a name.")
		return false
	}
	if prefix == "" {
		Print("Gotools Error: function TakeRole() You need to declare a prefix.\nMake sure you have the prefix variable set in main() function.")
		return false
	}

	str := Replace(m.Content, prefix + Command.Color + " ", "", -1) // Remove the Prefix amd Command from the string.

	// Find out if they are using the command correctly.
	if Contains(str, " ") == false {
		// they're not using the command properly. let's help!
		SendMessage(s, m.ChannelID, ColorErrorResponse)
		return false
	}

	therole := Replace(m.Content, prefix + Command.Color + " " + Split(str, " ")[0] + " ", "", -1)
	thecolor := Split(str, " ")[0]
	Role = therole
	Color = thecolor


	guildID, err := ServerID(s, m)
	if err != nil {
		// can't find server.
		return false
	}
	thecolor = strings.Replace(thecolor, "#", "", -1)

	roles, err := s.State.Guild(guildID)
	if err == nil {
		for _, v := range roles.Roles {
    		if v.Name == therole {
    			roleID = v.ID
    			hoist = v.Hoist
    			perms = v.Permissions
    		}
    	}
	} else {
		//	fmt.Println("s.GuildRoles is the error")
	}

	var ij int
	newcode, _ := strconv.ParseInt(thecolor, 16, 0)
	d := fmt.Sprintf("%d", newcode)
	// fmt.Println(d)
	ij, err = strconv.Atoi(d)
	if err != nil {
		return false
	}

	_, err = s.GuildRoleEdit(guildID, roleID, therole, ij, hoist, perms)
	if err != nil {
		fmt.Println("Role Color Error:")
		fmt.Println(err)
		return false
	} else {
		return true
	}

	return false
}





func GiveRole(s *discordgo.Session, m *discordgo.MessageCreate, prefix string, cmd string) bool {

	if cmd == "" {
		Print("Gotools Error: function TakeRole() You need to declare a Command name\nUse gto.SetName.Take to set a name.")
		return false
	}
	if prefix == "" {
		Print("Gotools Error: function TakeRole() You need to declare a prefix.\nMake sure you have the prefix variable set in main() function.")
		return false
	}
	str := Replace(m.Content, prefix + cmd + " ", "", -1) // replaces the prefix and command name from the string.
	if Contains(str, " ") == false { // Look to see if they have the user id AND the role name.
		SendMessage(s, m.ChannelID, GiveErrorResponse)
		return false
	}

	if Contains(str, "<@") == false { // Let's make sure an ID Exists to prevent any errors.
		SendMessage(s, m.ChannelID, GiveErrorResponse)
		return false
	}


	// Now we need to clean this up a bit.
	data := Split(str, "<@") // first out of three steps to cleaning the id.
	a := Split(data[1], ">") // second step notice the data[1]
	therole := Replace(m.Content, prefix + cmd + " <@"+a[0]+"> ", "", -1) // clean the original string to get the Role Name.
	theuser := Replace(a[0], "!", "", -1) // nickname support. If they have a nickname it adds an !
	User = theuser
	Role = therole




	// server check
	guildID, err := ServerID(s, m)
	if err != nil {
		Print("Error Grabbing Server from Discord.")
		return false
	}

	// Print("Server ID: " + guildID)

	// member check
	x, err := s.State.Member(guildID, theuser) 
	if err != nil {
		Print("Error Grabbing User from Discord.")
		return false
	}


	// check if the server has the role.
	roles, err := s.State.Guild(guildID)
	if err == nil {
		for _, v := range roles.Roles {
    		if v.Name == therole {
    			x.Roles = append(x.Roles, v.ID)
    			s.GuildMemberEdit(guildID, theuser, x.Roles)
    			return true
    		}
		}
	}
	return false
}






func AutoRoleSystem(s *discordgo.Session, m *discordgo.GuildMemberAdd) bool {
	therole := AutoRoleName
	theuser := m.User.ID
	User = theuser
	Role = therole

	if AutoRoleBots == false && m.User.Bot == true {
		return false
	}

	// server check
	guildID := m.GuildID

	if err != nil {
		Print("Error Grabbing Server from Discord.")
		return false
	}

	// Print("Server ID: " + guildID)

	// member check
	x, err := s.State.Member(guildID, theuser) 
	if err != nil {
		Print("Error Grabbing User from Discord.")
		return false
	}



	// check if the server has the role.
	roles, err := s.State.Guild(guildID)
	if err == nil {
		for _, v := range roles.Roles {
    		if v.Name == therole {
    			x.Roles = append(x.Roles, v.ID)
    			s.GuildMemberEdit(guildID, theuser, x.Roles)
    			return true
    		}
		}
	}
	return false
}








func AddRole(s *discordgo.Session, m *discordgo.MessageCreate, therole string, permissions int) bool {

	// if they don't know what to do with this they can leave it to 0
	// I will convert to "Basic Permissions" for a new channel.
	if permissions == 0 {
		permissions = 11656193
	}

	guildID, err := ServerID(s, m)
	if err != nil {
		// couldn't find the server.
		return false
	}

	if GetRoleID(s, guildID, therole) == "" {
		com, err := s.GuildRoleCreate(guildID)
		if err != nil {
			Print("Gotools function AddRole() Error:")
			Print(err)
			return false
		}

		_, err = s.GuildRoleEdit(guildID, com.ID, therole, 0, false, permissions)
		if err != nil {
			Print("Gotools function AddRole() Error:")
			Print(err)
			return false
		}

		return true
	} else {
		// the role already exists.
		return false
	}
}







func DeleteRole(s *discordgo.Session, m *discordgo.MessageCreate, therole string) bool {
	// server check
	guildID, err := ServerID(s, m)
	if err != nil {
		// couldn't find the server.
		return false
	}

	// role check
	roleid := GetRoleID(s, guildID, therole)
	if roleid != "" {
		err = s.GuildRoleDelete(guildID, roleid)
		if err != nil {
			// error editing roles. usually means permissions!
			return false
		}
		// it worked!
		return true
	} else {
		// role doesn't exist
		return false
	}
}














func InVoiceChannel(s *discordgo.Session, guildID string, user string) bool {
g, err := s.State.Guild(guildID)
if err != nil {
	// Could not find guild.
	return false
}
cnt := 0
for _, vs := range g.VoiceStates {
	if vs.UserID == user {
		cnt++
	}

	if cnt == 0 {
		return false
	}
	if cnt > 0 {
		return true
	}
}
return false
}
