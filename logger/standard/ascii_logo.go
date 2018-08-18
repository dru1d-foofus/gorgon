package standard

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/dru1d-foofus/gorgon"
)

// PrintLogo prints the genesis logo to the current standard out
func PrintLogo() {
	fmt.Fprintf(color.Output, "%s\n", ASCIILogo())
}

// ASCIILogo returns a string containing a color formatted logo block
// ASCII art sourced here: http://ascii.co.uk/art/medusa by a user named jgs
func ASCIILogo() string {
	lines := []string{
		errorLevel.Sprint("************************************************"),
		infoMsg.Sprint("                   ,--.                            "),
		infoMsg.Sprint("          ,--.  .--,`) )  .--,                     "),
		infoMsg.Sprint("       .--,`) \\( (` /,--./ (`                      "),
		infoMsg.Sprint("      ( ( ,--.  ) )\\ /`) ).--,-.                   "),
		infoMsg.Sprint("       ;.__`) )/ /) ) ( (( (`_) )                  "),
		infoMsg.Sprint("      ( (  / /( (.' \"-.) )) )__.'-,                "),
		infoMsg.Sprint("     _,--.( ( /`         `,/ ,--,) )               "),
		infoMsg.Sprint("    ( (``) \\,` ==.    .==  \\( (`,-;                "),
		infoMsg.Sprintf("     ;-,( (_) ~%s~ \\  / ~%s~ (_) )_) )               ", color.HiRedString("6"), color.HiRedString("6")),
		infoMsg.Sprint("    ( (_ \\_ (      )(      )__/___.'               "),
		infoMsg.Sprint("    '.__,-,\\ \\     ''     /\\ ,-.                   "),
		infoMsg.Sprint("       ( (_/ /\\    __    /\\ \\_) )                  "),
		infoMsg.Sprint("        '._.'  \\  \\__/  /  '._.'                   "),
		infoMsg.Sprint("            .--`\\      /`--.                       "),
		infoMsg.Sprint("                 '----'                            "),
		infoMsg.Sprint("                                                   "),
		fmt.Sprintf("      %s   %s    ", infoLevel.Sprint("GOrgon"), defaultLevel.Sprintf("v%s", gorgon.Version)),
		fmt.Sprintf("	  	        %s   \n                %s    ", color.HiGreenString("-- By --"), color.HiRedString("dru1d")),
		"          github.com/dru1d-foofus/gorgon",
		errorLevel.Sprint("************************************************"),
	}

	return strings.Join(lines, "\n")
}
