package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

const logo = `
/______/\  /_____/\ /_____/\ /__/\ /__/\ /_____/\ /_____/\   /_______/\  /_____/\/_____/\     
\::::__\/__\:::_ \ \\:::_ \ \\::\ \\  \ \\::::_\/_\:::_ \ \  \__.::._\/  \:::__\/\::::_\/_    
 \:\ /____/\\:\ \ \ \\:(_) \ \\::\/_\ .\ \\:\/___/\\:(_) ) )_   \::\ \      /: /  \:\/___/\   
  \:\\_  _\/ \:\ \ \ \\: ___\/ \:: ___::\ \\::___\/_\: __ \\ \  _\::\ \__  /::/___ \::___\/_  
   \:\_\ \ \  \:\_\ \ \\ \ \    \: \ \\::\ \\:\____/\\ \ \\ \ \/__\::\__/\/_:/____/\\:\____/\ 
    \_____\/   \_____\/ \_\/     \__\/ \::\/ \_____\/ \_\/ \_\/\________\/\_______\/ \_____\/ 
    `
const description = `
    Thank you for installing gopherize!!

    Welcome to Gopherize, your go-to resource for learning Go! This CLI will guide you through
    a series of exercises designed to teach you the fundamentals of Go programming. Each exercise
    is designed to be simple and easy to follow, even if you're a complete beginner.You'll 
    start by learning the basics of the language and then gradually move on to more advanced topics.
    By the time you've completed all the exercises, you'll have a solid foundation in Go and be 
    ready to tackle more complex projects. So let's get started and Gopherize your coding skills!
    `

func Root() *cobra.Command {
	b := color.FgBlue.Render
	var rootCmd = &cobra.Command{
		Use:   "gopherize <command>",
		Short: "Let's gopherize you into gopher!",
		Long:  b(logo) + "\n" + description,
	}

	rootCmd.AddCommand(runCmd())
	rootCmd.AddCommand(verifyCmd())
	rootCmd.AddCommand(watchCmd())
	return rootCmd
}
