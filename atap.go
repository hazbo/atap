/**
 * atap v0.1-dev
 *
 * (c) Ground Six
 *
 * @package atap
 * @version 0.1-dev
 * 
 * @author Harry Lawrence <http://github.com/hazbo>
 *
 * License: MIT
 * 
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package main

import (
    "./vendor/jconfig"
    "fmt"
    "log"
    "os"
    "os/exec"
)

/**
 * @var string name of the command
 * @var int the number of times to run
 * @args []string command arguments
 */
type Command struct {
    name  string
    count int
    args  []string
}

/**
 * Entry point for atap. Runs the command
 * and follows general instructions as
 * specefied in the JSON configuration
 * file
 */
func main() {
    configFile := getConfigFile()
    if configFile != "" {
        command := new(Command)

        config := jconfig.LoadConfig(getConfigFile())
        command.name = config.GetString("command")
        command.count = config.GetInt("count")

        args := config.GetArray("args")
        argsSlice := make([]string, len(args))

        for i, s := range args {
            argsSlice[i] = s.(string)
        }

        command.args = argsSlice

        executeCommand(command)
    }
}

/**
 * Gets the name of your config file
 * from the param passed through when
 * the program is ran
 *
 * e.g. atap config.json
 *
 * @return string file path or empty string
 */
func getConfigFile() string {
    if len(os.Args) > 1 {
        filePath := os.Args[1]
        if _, err := os.Stat(filePath); err == nil {
            return filePath
        }
    }
    return ""
}

/**
 * Executes command a given amount
 * of times as specefied in the
 * JSON configuration file
 *
 * @param Command in instance of Command struct
 *
 * @return nil
 */
func executeCommand(command *Command) {
    for i := 0; i < command.count; i++ {
        out, err := exec.Command(command.name, command.args...).Output()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%s\n", out)
    }
}
