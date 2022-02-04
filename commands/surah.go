package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/df-mc/dragonfly/server/world/sound"
)

type Surah struct {
	Surah int32
}

func (c Surah) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)
	if !ok {
		return // run ingame, allah can only be summoned by the grace of a human character
	}

	link := fmt.Sprintf("https://raw.githubusercontent.com/semarketir/quranjson/master/source/surah/surah_%d.json", c.Surah)
	res, err := http.Get(link)
	if err != nil {
		p.Message(err.Error())
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		p.Message(err.Error())
		return
	}

	var surah map[string]interface{}
	err = json.Unmarshal(body, &surah)
	if err != nil {
		p.Message(err.Error())
		return
	}

	verses := surah["verse"].(map[string]interface{})
	i := 1

	t := title.New(fmt.Sprintf("§4§lSurah %d: %s",  c.Surah, surah["name"]))
	p.SendTitle(t)

	p.World().PlaySound(p.Position(), sound.ChestOpen{})
	p.World().PlaySound(p.Position(), sound.Thunder{})

	go func() {
		for _, verse := range verses {
			msg := fmt.Sprintf("§4§l %d > %s", i, verse)
			p.Message(fmt.Sprintln(msg))
			i++
		}
	}()

}