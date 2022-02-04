package commands

import "github.com/df-mc/dragonfly/server/cmd"

func Register() {
	cmd.Register(cmd.New("callallah", "Call the mighty Allah 'salaa allah ealayh wasalam'", []string{"allah"}, CallAllah{}))
	cmd.Register(cmd.New("surah", "Call to the Holy Quran's Surah's", []string{}, Surah{}))

}