package rdb

import (
	"encoding/json"
	"fmt"
	"github.com/fzzy/radix/redis"
	ferr "github.com/jonas747/fortia/error"
	"github.com/jonas747/fortia/messages"
	"github.com/jonas747/fortia/vec"
	"github.com/jonas747/fortia/world"
	"strconv"
	"sync"
)

// Basic implementation of world.GameDB
type GameDB struct {
	*Database
}

// Information about the world. Blocktypes, biomes, layer size  and such
func (g *GameDB) GetWorldInfo() (*world.WorldInfo, ferr.FortiaError) {
	var info world.WorldInfo
	err := g.GetJson("worldinfo", &info)
	return &info, err
}

func (g *GameDB) SetWorldInfo(info *world.WorldInfo) ferr.FortiaError {
	return g.SetJson("worldinfo", info)
}

// Returns the specified user's info
func (g *GameDB) GetUserInfo(user string) (*world.UserInfo, ferr.FortiaError) {
	infoHash, err := g.GetHash("u:" + user)
	if err != nil {
		return nil, err
	}
	return &world.UserInfo{
		Name: infoHash["name"],
	}, nil
}

// Sets the specified users info fields from info map to whatever is in the info map
func (g *GameDB) SetUserInfo(info *world.UserInfo) ferr.FortiaError {
	infoHash := map[string]interface{}{
		"name": info.Name,
	}
	return g.SetHash("u:"+info.Name, infoHash)
}

func (g *GameDB) GetUserEntities(user string) ([]int, ferr.FortiaError) {
	reply, err := g.Cmd("SMEMBERS", "ue:"+user)
	if err != nil {
		return []int{}, err
	}

	list, nErr := reply.List()
	if nErr != nil {
		return []int{}, ferr.Wrap(nErr, "")
	}

	intSlice := make([]int, len(list))
	for k, v := range list {
		intSlice[k], _ = strconv.Atoi(v)
	}
	return intSlice, nil
}

func (g *GameDB) EditUserEntities(user string, add, del []int) ferr.FortiaError {
	return g.EditSetI("ue:"+user, add, del)
}

// Returns the chunkinfo
func (g *GameDB) GetChunk(pos vec.Vec2I) (*messages.Chunk, ferr.FortiaError) {
	var chunk messages.Chunk
	err := g.GetProto(fmt.Sprintf("c:%d:%d", pos.X, pos.Y), &chunk)
	return &chunk, err
}

func (g *GameDB) SetChunk(chunk *messages.Chunk) ferr.FortiaError {
	return g.SetProto(fmt.Sprintf("c:%d:%d", chunk.X, chunk.Y), chunk)
}

// Get and set entities
func (g *GameDB) GetEntity(id int) (map[string]string, ferr.FortiaError) {
	return g.GetHash("e:" + strconv.Itoa(id))
}

func (g *GameDB) SetEntity(id int, info map[string]interface{}) ferr.FortiaError {
	return g.SetHash("e:"+strconv.Itoa(id), info)
}

func (g *GameDB) PopAction(tick int) (*world.Action, ferr.FortiaError) {
	reply, err := g.Cmd("SPOP", fmt.Sprintf("actionQueue:%d", tick))
	if err != nil {
		return nil, err
	}

	if reply.Type == redis.NilReply {
		// Not found
		return nil, ferr.Newc("No Actions for that tick number", world.ErrCodeNotFound)
	}

	raw, nErr := reply.Bytes()
	if nErr != nil {
		return nil, ferr.Newc("Error converting", world.ErrCodeConversionErr)
	}

	var action *world.Action
	nErr = json.Unmarshal(raw, action)
	if nErr != nil {
		return nil, ferr.Wrap(nErr, "")
	}

	return action, nil
}

func (g *GameDB) IncrTick() (int, ferr.FortiaError) {
	reply, err := g.Cmd("INCR", "tick")
	if err != nil {
		return 0, err
	}

	n, nErr := reply.Int()
	if nErr != nil {
		return n, ferr.Wrapc(nErr, world.ErrCodeConversionErr)
	}
	return n, nil
}
