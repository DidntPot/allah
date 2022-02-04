package commands

import (
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/bossbar"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/df-mc/dragonfly/server/world/sound"
)

type CallAllah struct {}

func (c CallAllah) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)
	if !ok {
		return // run ingame, allah can only be summoned by the grace of a human character
	}

	l := entity.NewLightning(p.Position())
	p.World().AddEntity(l)

	p.Message("§4§lThe Almighty ALLAH (صلى الله عليه وسلم) HAS ARRIVED!")

	bb := bossbar.New("§4§lبِسْمِ ٱللَّهِ ٱلرَّحْمَٰنِ ٱلرَّحِيمِ")
	p.SendBossBar(bb)

	t := title.New("§4§lٱلْحَمْدُ لِلَّهِ رَبِّ ٱلْعَٰلَمِينَ")
	p.SendTitle(t)

	go func() {
		for i := 0; i < 3; i++ {
			p.World().PlaySound(p.Position(), sound.DoorCrash{})
			p.World().PlaySound(p.Position(), sound.Pop{})
			p.World().PlaySound(p.Position(), sound.Explosion{})
		}	
	}()

	go func() {
		time.Sleep(5 * time.Second)
		p.RemoveBossBar()
	}()

}