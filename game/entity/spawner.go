package entity

import (
	"math/rand"

	"github.com/conejoninja/gopherbadge/game/assets"
)

type EnemySpawner struct {
	CheckFrequencyInSeconds float32
	ChanceOfSpawning        float32
	distanceSinceLastSpawn  float32
}

func NewEnemySpawner(chanceOfSpawning float32) *EnemySpawner {
	return &EnemySpawner{
		ChanceOfSpawning:       chanceOfSpawning,
		distanceSinceLastSpawn: 0,
	}
}

func (e *EnemySpawner) SpawnEnemy(distanceTraveled float32) *EnemyEntity {
	e.distanceSinceLastSpawn += distanceTraveled

	if e.distanceSinceLastSpawn >= EnemySpawnDistance {
		diceResult := rand.Float32()
		if diceResult <= e.ChanceOfSpawning {
			e.distanceSinceLastSpawn = 0
			return &EnemyEntity{
				Entity:        NewEntity(320, 160, 40, 60, assets.Bug1), //TODO replace with variables instead of magic numbers
				DidCollide:    false,
				HasBeenScored: false,
			}
		}
	}
	return nil
}
