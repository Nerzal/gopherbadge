package entity

import (
	"math"
	"math/rand"
)

type EnemySpawner struct {
	CheckFrequencyInSeconds float32
	ChanceOfSpawning        float32
	distanceSinceLastSpawn  float32
}

func NewEnemySpawner(checkFrequencyInSeconds, chanceOfSpawning float32) *EnemySpawner {
	return &EnemySpawner{
		CheckFrequencyInSeconds: checkFrequencyInSeconds,
		ChanceOfSpawning:        chanceOfSpawning,
		distanceSinceLastSpawn:  0,
	}
}

func (e *EnemySpawner) SpawnEnemy(distanceTraveled float32) *EnemyEntity {
	e.distanceSinceLastSpawn += distanceTraveled

	if e.distanceSinceLastSpawn >= EnemySpawnDistance && rand.Float32()/math.MaxFloat32 >= e.ChanceOfSpawning {
		return &EnemyEntity{
			Entity:        NewEntity(160, 320, 80, 50), //TODO replace with variables instead of magic numbers
			DidCollide:    false,
			HasBeenScored: false,
		}
	}
	return nil
}
