package entity_test

import (
	"testing"

	"github.com/conejoninja/gopherbadge/game/entity"
)

func TestHasCollision(t *testing.T) {
	tt := []struct {
		description   string
		player        *entity.Entity
		enemy         *entity.Entity
		shouldCollide bool
	}{
		{
			description: "player and enemy overlap completely",
			player: &entity.Entity{
				PosX:   10,
				PosY:   8,
				Width:  20,
				Height: 20,
			},
			enemy: &entity.Entity{
				PosX:   10,
				PosY:   8,
				Width:  20,
				Height: 20,
			},
			shouldCollide: true,
		},
		{
			description: "player and enemy overlap on x axis",
			player: &entity.Entity{
				PosX:   10,
				PosY:   8,
				Width:  20,
				Height: 20,
			},
			enemy: &entity.Entity{
				PosX:   0,
				PosY:   8,
				Width:  20,
				Height: 20,
			},
			shouldCollide: true,
		},
		{
			description: "player and enemy overlap on x axis but not on y axis",
			player: &entity.Entity{
				PosX:   10,
				PosY:   29,
				Width:  20,
				Height: 20,
			},
			enemy: &entity.Entity{
				PosX:   0,
				PosY:   8,
				Width:  20,
				Height: 20,
			},
			shouldCollide: false,
		},
		{
			description: "player and enemy overlap on x axis and on y axis",
			player: &entity.Entity{
				PosX:   10,
				PosY:   8,
				Width:  20,
				Height: 20,
			},
			enemy: &entity.Entity{
				PosX:   0,
				PosY:   28,
				Width:  20,
				Height: 20,
			},
			shouldCollide: true,
		},
		{
			description: "player and enemy do not overlap at all",
			player: &entity.Entity{
				PosX:   10,
				PosY:   8,
				Width:  20,
				Height: 20,
			},
			enemy: &entity.Entity{
				PosX:   100,
				PosY:   28,
				Width:  20,
				Height: 20,
			},
			shouldCollide: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			collision := tc.player.HasCollision(tc.enemy)
			if collision != tc.shouldCollide {
				t.Errorf("expected %v, got %v", tc.shouldCollide, collision)
			}
		})
	}
}
