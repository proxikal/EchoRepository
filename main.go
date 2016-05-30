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
	"strings"
	"errors"
	"strconv"
	//"log"
	"github.com/bwmarrin/discordgo"
)



type error interface {
    Error() string
}

type errorString struct {
    s string
}


type pushString struct {
	s string
}


func (e *errorString) Error() string {
    return e.s
}

func New(text string) error {
    return &errorString{text}
}


	var err error


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





func getJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}






func readJson(path string, key string) (interface{}, error) {
//	var newval map[string]interface{}
	var rjson map[string]interface{}

	// ################## COMMANDS.JSON UPDATE ########################
	file, err := ioutil.ReadFile(path)
	if err != nil {
	//	fmt.Println("gotools error =>")
	//	fmt.Println(err)
		return "", err
	} else {
		json.Unmarshal(file, &rjson)
	//	for k1, v1 := range rjson {
		if _, ok := rjson[key]; ok {
			/*
			switch rjson[key].(type) {
				case int:
					return rjson[key].(int), nil
				case string:
					return rjson[key].(string), nil
				case float64:
					newval = rjson[key].(float64)
				case int32:
					newval = rjson[key].(int32)
				case int64:
					newval = rjson[key].(int64)
				}
				*/

		return rjson[key].(interface{}), nil
		//	 	newval = v1.(type)
		}
	//	} // end of oldcommand for loop
	} // check if error is nil
	return nil, err
}



func writeJson(path string, key string, val string) error {
//	newval := ""
	var rjson map[string]interface{}

	// ################## COMMANDS.JSON UPDATE ########################
	file, err := ioutil.ReadFile(path)
	if err != nil {
	//	fmt.Println("gotools error =>")
	//	fmt.Println(err)
		return err
	} else {
		json.Unmarshal(file, &rjson)
		rjson[key] = val
	b, err := json.MarshalIndent(rjson, "", "   ")
	if err == nil {
		ioutil.WriteFile(path, b, 0777)
		// it works
	}	else {
		return err
	}

	} // check if error is nil
return err
}




func ReadFile(path string) ([]byte, error) {
	tb, err := ioutil.ReadFile(path)
	return tb, err
}


func WriteFile(path string, b []byte, perm os.FileMode) {
	ioutil.WriteFile(path, b, perm)
}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}



func Split(str string, op string) []string {
	return strings.Split(str, op)
}



func readLines(path string) ([]string, error) {
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




func countLines(path string) int {
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






// ################ DiscordGo Manipulation Below #################


func SendMessage(s *discordgo.Session, channelID string, text string) {
	s.ChannelMessageSend(channelID, text)
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



/*
type Info struct {
	Servers		int
	Roles 		int
	Channels 	int
	Members 	int
}
*/



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
