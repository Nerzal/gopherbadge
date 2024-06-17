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
