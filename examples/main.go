package main


import (
	"gotools"
	"net/url"
	"time"
	"github.com/bwmarrin/discordgo"
)


type config struct {
	Token					string
	Servers					int
	OwnerID					string
	TwitterConsumerKey		string
	TwitterConsumerSecret	string
	TwitterToken			string
	TwitterTokenSecret		string
	RiotKey					string
	Greet 					bool
	GreetMessage 			string
	PmGreet 				bool
	PmGreetMessage			string
	AutoRole 				bool
	AutoRoleName 			string
}
	var conf config
	var Prefix string



func loadConf() {
	// grabbing config.json file
	file, err := gto.ReadFile("config.json")
	if err != nil {
		gto.Print("Can't find config.json file. Make sure it's in the same path as the bot.")
		return
	}
	// loading the json from the file into the config structure.
	gto.Unmarshal(file, &conf)
}







func Debug(s *discordgo.Session, channelID string) {
	dbg:
	loadConf()

	gto.SendMessage(s, channelID, "**Gotools Debug System**")
	time.Sleep(10 * time.Minute)

	gto.SendMessage(s, channelID, "```AutoIt\n@gto.Autorole = \""+conf.AutoRole+"\"\n@gto.AutoRoleName = \""+conf.AutoRoleName+"\"\n@gto.Greet = \""+conf.Greet+"\"\n@gto.GreetMessage = \""+conf.GreetMessage+"\"\n@gto.PmGreet = \""+conf.PmGreet+"\"\n@gto.PmGreetMessage = \""+conf.PmGreetMessage+"\"```")
	After(5 * time.Minute)
	goto dbg
}












func main() {
	// grabbing config.json file
	loadConf()



	// let's load some stuff in.
	// Change the role to whatever role you want to control the bot.
	gto.Master = "Bot Commander"



	// let's set the bots prefix. You can change this later on.
	Prefix = "-"



	// Let's name the commands.
	gto.Command.Give = "give"
	gto.Command.Take = "take"
	gto.Command.Color = "rolecolor"



	// Let's set what Gotools will say when an error is triggered in Give and Take Commands.
	gto.GiveErrorResponse = "An example of this command: `"+Prefix+gto.Command.Give+" @User Role Name`" // set what gotools will say when there is an error.
	gto.TakeErrorResponse = "An example of this command: `"+Prefix+gto.Command.Take+" @User Role Name`" // set what gotools will say when there is an error.
	gto.ColorErrorResponse = "An example of this command: `"+Prefix+gto.Command.Color+" #FF0000 Role Name`"
	// REMEMBER: if you change the command names change them in the above message ^
	// you can change any command name in the function of the command.






	// connecting to discord.
	dg, err := gto.Connect(conf.Token)
	if err != nil {
		gto.Print("Wasn't able to connect to Discord. Check your token and try again!")
		return
	}



	// Register events
	dg.AddHandler(messageCreate)
	dg.AddHandler(onReady)
	dg.AddHandler(GuildMemberAdd)
	dg.AddHandler(GuildMemberRemove)
	// dg.AddHandler(GuildCreate)
	// dg.AddHandler(GuildDelete)
	dg.Open()

	gto.Print("Running Gotools Version 1.0")
	// keeps the program running until you hit a key in the terminal.
	gto.Listen()
}




// Discord will send a ready packet when the bot has logged in.
// this function is registered above using Addhandler
func onReady(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Gotools 1.0.0")
	gto.Print("Ready Packet recieved.")
}




// what happens when someone joins your server.
// this function is registered above using Addhandler
func GuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

	// let's parse a few keys.
	gto.GreetMessage = gto.Replace(gto.GreetMessage, "{user}", "<@"+m.User.ID+">", -1)
	gto.PmGreetMessage = gto.Replace(gto.PmGreetMessage, "{user}", "<@"+m.User.ID+">", -1)



	gto.AutoRoleListen(s, m)
	if gto.Greet == true {
		gto.SendMessage(s, m.GuildID, gto.GreetMessage)
	}

	if gto.PmGreet == true {
		pm, err := s.UserChannelCreate(m.User.ID)
		if err != nil {
			// can't pm due to blocked or private.
			return
		} else {
			gto.SendMessage(s, pm.ID, gto.PmGreetMessage)
		}
	}
}




// what happens when someone leaves your server.
// this function is registered above using Addhandler
func GuildMemberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {

}



// This function will trigger everytime someone sends a message (that your bot can see)
// this function is registered above using Addhandler
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {


	loadConf()
	// Early beta below. this may not work (still needs testing)
	gto.Greet = conf.Greet
	gto.AutoRole = conf.AutoRole
	gto.AutoRoleName = conf.AutoRoleName
	gto.GreetMessage = conf.GreetMessage
	gto.PmGreet = conf.PmGreet
	gto.PmGreetMessage = conf.PmGreetMessage





	cid, err := gto.ServerID(s, m)
	if err != nil {
		// had issues grabbing the server id.
		return
	}




	server, err := s.State.Guild(cid)
	if err != nil {
		// Could not find guild.
		return
	}





	// let's set up a master system. By default the role will be Bot Commander
	isMaster := false
	if gto.MemberHasRole(s, server.ID, m.Author.ID, gto.Master) {
		isMaster = true
	}




	// Detect the server owner and make him Master as well..
	if m.Author.ID == server.OwnerID {
		isMaster = true
	}






	// let's create a basic help command to let people know what commands I have.
	if gto.HasPrefix(m.Content, Prefix + "help") {
		gto.SendMessage(
			s, m.ChannelID, "```ruby\nBasic Bot Example\n"+Prefix+"help\n"+Prefix+"give @User Role Name\n"+Prefix+"take @User Role Name\n"+Prefix+"rolecolor #C0C0C0 Role Color\n"+Prefix+"addrole Role Name\n"+Prefix+"delrole Role Name```")
	}






	// Bot INTERNAL COMMANDS Change Name & Change Avatar
	// let's set the -avatar command to change the bots picture.
	if gto.HasPrefix(m.Content, Prefix + "avatar") {
		path := "new.png" // this is assuming the file new.png is in the bots directory.
		if m.Author.ID == conf.OwnerID {
			if gto.ChangeAvatar(s, path) {
				// it worked!
				gto.SendMessage(s, m.ChannelID, "I have changed my avatar.")
			} else {
				// an error occured.
				gto.SendMessage(s, m.ChannelID, "I wasn't able to complete the task.")
			}
		} // Only the bot owner can do this.
	}








	if gto.HasPrefix(m.Content, Prefix + "name ") {
		if isMaster == true {
			str := gto.Replace(m.Content, Prefix + "name ", "", -1)
			if gto.ChangeName(s, str) == true {
				gto.SendMessage(s, m.ChannelID, "I have changed my name.")
			} else {
				gto.SendMessage(s, m.ChannelID, "There was an error trying to change the name.")
			}
		}
	}





	if gto.HasPrefix(m.Content, Prefix + "greet ") {
		str := gto.Replace(m.Content, Prefix + "greet ", "", -1)

		if str == "off" {
			conf.Greet = false
			j, err := gto.Marshal(conf)
			if err != nil {
				// json marshal error
				return
			}
			gto.WriteFile("config.json", j, 0777)
		}


		if str != "" && str != "off" {
			conf.Greet = true
			conf.GreetMessage = str
			j, err := gto.Marshal(conf)
			if err != nil {
				// json marshal error
				return
			}
			gto.WriteFile("config.json", j, 0777)
		}
	}








	if gto.HasPrefix(m.Content, Prefix + "pm ") {
		str := gto.Replace(m.Content, Prefix + "pm ", "", -1)

		if str == "off" {
			conf.PmGreet = false
			j, err := gto.Marshal(conf)
			if err != nil {
				// json marshal error
				return
			}
			gto.WriteFile("config.json", j, 0777)
		}


		if str != "" && str != "off" {
			conf.PmGreet = true
			conf.PmGreetMessage = str
			j, err := gto.Marshal(conf)
			if err != nil {
				// json marshal error
				return
			}
			gto.WriteFile("config.json", j, 0777)
		}
	}







	if gto.HasPrefix(m.Content, Prefix + "image") {
		go gto.PostRandomImage(s, m, "System/images")
	}















	// Twitter Commands Example:
	// Let people tweet on your bot's twitter live example: (http://twitter.com/echothebot)
	if gto.HasPrefix(m.Content, Prefix + "tweet ") {
		str := gto.Replace(m.Content, Prefix + "tweet ", "", -1)
		if str == "" {
			gto.SendMessage(s, m.ChannelID, "An example of this command `"+Prefix+"tweet Your message here`")
			return
		}
		// let's set the Twitter information using Anaconda's system through Gotools.
		gto.ConsumerKey(conf.TwitterConsumerKey)
    	gto.ConsumerSecret(conf.TwitterConsumerSecret)
		api := gto.TwitterApi(conf.TwitterToken, conf.TwitterTokenSecret)
		api.PostTweet(str, url.Values{"status": {str}})
		// api.PostTweet(message, url)
		gto.SendMessage(s, m.ChannelID, "I have posted your message on my twitter!")
	}






	if gto.HasPrefix(m.Content, Prefix + "riot coop ") {

	}










	// Give someone a role.
	// Let's make this command Masters only (Server Owner & Bot Commanders)
	if gto.HasPrefix(m.Content, Prefix + gto.Command.Give + " ") {
		if isMaster == true { // Check if the user is a master or not.
			if gto.GiveRole(s, m, Prefix, gto.Command.Give)  == true { // we're going to send the prefix & the command name we want (you can change this)
				// for example if you want this command to be called GiveIt you do gto.GiveRole(s, m, prefix, "GiveIt")
				// just be sure to change the gto.HasPrefix() line to account for the new command name.
				// it worked!
				// gto.SendMessage(s, m.ChannelID, "I have given the user <@"+user+"> the role `"+role+"`")
				gto.SendMessage(s, m.ChannelID, "I have given <@"+gto.User+"> the role `"+gto.Role+"`")
			} else {
				// this usually means permission problems. Let's let them know
				// Also discord changed the role system around with the Administrator entry. This requires Gotools rank to be higher
				// than theirs in order to edit them.
				gto.SendMessage(s, m.ChannelID, "Something went wrong. *(Usually Permissions)* Check to make sure I have the appropriate permissions")
			}
		} else { // the user is not my master.
			gto.SendMessage(s, m.ChannelID, "You're not a Bot Commander.")
		}
	} // end of -give @User Role Name command.















	// Take a role away from someone.
	// Let's make this command Masters only (Server Owner & Bot Commanders)
	// this command is basically the same as -give we just need to tweak a few things here and there.
	if gto.HasPrefix(m.Content, Prefix + gto.Command.Take + " ") {
		if isMaster == true {
			if gto.TakeRole(s, m, Prefix, gto.Command.Take) == true {
				// it worked!
				gto.SendMessage(s, m.ChannelID, "I have taken the role `"+gto.Role+"` from <@"+gto.User+">")
			} else {
				// an error has occured. (usually permission errors)
				gto.SendMessage(s, m.ChannelID, "An error occured. Check my permissions *(Drag my role to the top to manage Admins rols)*")
			}
		} else { // the user is not my master.
			gto.SendMessage(s, m.ChannelID, "You're not a Bot Commander.")
		}
	} // end of -take @User Role Name command.















	// Let's make the role color command. Again thanks to Gotools all the heavy stuff is already done.
	// we just need to parse the message to grab the role name, and the command.
	if gto.HasPrefix(m.Content, Prefix + gto.Command.Color + " ") {
		if isMaster == true {
			if gto.RoleColor(s, m, Prefix, gto.Command.Color) { // Call to Gotools function RoleColor()
				// it worked!
				gto.SendMessage(s, m.ChannelID, "I have changed `"+gto.Role+"` color to `"+gto.Color+"`")
			}
		} else {
			// an error occured (usually permissions) or the role doesn't exist.
			gto.SendMessage(s, m.ChannelID, "An error occured. Check my permissions. Make sure the role exists.")
			return
		}

	}














	// Let's make the -addrole command
	// Again all we have to do is parse the message. Gotools will do the rest!
	if gto.HasPrefix(m.Content, Prefix + "addrole ") {
		role := gto.Replace(m.Content, Prefix + "addrole ", "", -1) // parse the prefix and command out of the role.
		if role != "" {
			// they need to choose a role.
			gto.SendMessage(s, m.ChannelID, "You need to pick a role name `"+Prefix+"addrole Role Name`")
			return
		}

		if isMaster == true {
			if gto.AddRole(s, m, role, 0) { // you can change 0 to whatever permission bit integer or use 0 = 11656193
				// it worked!
				gto.SendMessage(s, m.ChannelID, "I have added the role `"+role+"`")
			} else {
				// an error occured (usually permissions)
				gto.SendMessage(s, m.ChannelID, "Something happened. Check my permissions!")
				return
			}
		} else {
			// the user is not a master.
			gto.SendMessage(s, m.ChannelID, "Only Bot Commanders can use this command.")
		}
	}











	// Let's make the -delrole command
	// Again all we have to do is parse the message. Gotools will do the rest!
	if gto.HasPrefix(m.Content, Prefix + "delrole ") {
		role := gto.Replace(m.Content, Prefix + "delrole ", "", -1) // parse the prefix and command out of the role.
		if role != "" {
			// they need to choose a role.
			gto.SendMessage(s, m.ChannelID, "You need to pick a role name `"+Prefix+"delrole Role Name`")
			return
		}

		if isMaster == true {
			if gto.DeleteRole(s, m, role) {
				// it worked!
				gto.SendMessage(s, m.ChannelID, "I have removed the role `"+role+"`")
			} else {
				// an error occured (usually permissions)
				gto.SendMessage(s, m.ChannelID, "Something happened. Check my permissions!")
				return
			}
		} else {
			// the user is not a master.
			gto.SendMessage(s, m.ChannelID, "Only Bot Commanders can use this command.")
		}
	}












} // end of messageCreate Function.