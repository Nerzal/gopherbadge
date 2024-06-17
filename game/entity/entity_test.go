package entity_test

import (
	"github.com/conejoninja/gopherbadge/game/entity"
	"math"
	"testing"
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

func TestMove(t *testing.T) {
	tt := []struct {
		description       string
		enemies           []*entity.EnemyEntity
		timeDelta         float32
		enemySpeeds       []float32
		expectedPositions []float32
	}{
		{
			description: "enemies don't move if no time passes",
			enemies: []*entity.EnemyEntity{
				{
					Entity: &entity.Entity{
						PosX: 0,
					},
				},
				{
					Entity: &entity.Entity{
						PosX: 100,
					},
				},
			},
			timeDelta:         0,
			enemySpeeds:       []float32{12, 24},
			expectedPositions: []float32{0, 100},
		},
		{
			description: "enemies with a speed of 0 never move",
			enemies: []*entity.EnemyEntity{
				{
					Entity: &entity.Entity{
						PosX: 0,
					},
				},
			},
			timeDelta:         math.MaxFloat32,
			enemySpeeds:       []float32{0},
			expectedPositions: []float32{0},
		},
		{
			description: "enemies with a speed greater 0 move for in a time greater 0",
			enemies: []*entity.EnemyEntity{
				{
					Entity: &entity.Entity{
						PosX: 10,
					},
				},
				{
					Entity: &entity.Entity{
						PosX: 100,
					},
				},
			},
			timeDelta:         0.5,
			enemySpeeds:       []float32{10, 100},
			expectedPositions: []float32{5, 50},
		},
		{
			description: "enemies can move off screen (into the negative posX)",
			enemies: []*entity.EnemyEntity{
				{
					Entity: &entity.Entity{
						PosX: 10,
					},
				},
			},
			timeDelta:         0.5,
			enemySpeeds:       []float32{40},
			expectedPositions: []float32{-10},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			for idx := range tc.enemies {
				tc.enemies[idx].Move(tc.timeDelta, tc.enemySpeeds[idx])
				if tc.enemies[idx].PosX != tc.expectedPositions[idx] {
					t.Errorf("expected enemy at position %f, got %f", tc.expectedPositions[idx], tc.enemies[idx].PosX)
				}
			}
		})
	}
}
