package block

import (
    "time"
    "crypto/sha256"
	"encoding/hex"
	"encoding/json"
    "strconv"
    "strings"
    "fmt"
    "errors"
)

type Hash [32]byte

func (h *Hash) Hex() string {
	return hex.EncodeToString(h[:])
}

func hashS(params ...string) Hash {
    return sha256.Sum256([]byte(strings.Join(params[:], ",")))
}

type Data map[string]interface {}

func (d *Data) toSting() string {
    dat := ""
    for _, value := range *d {dat += fmt.Sprint(value, "\n")}
    return dat
}

type Time int64

func (t *Time) toSting() string {
    return strconv.FormatInt(int64(*t), 10)
}

func (t *Time) toTime() time.Time {
    return time.Unix(int64(*t), 0)
}

type Block struct {
    header string
    timestamp Time
    prevHash Hash
    hash Hash
    data Data
}

type block struct {
    Header string `json:"header"`
    Timestamp Time `json:"timestamp"`
    PrevHash Hash `json:"prev_hash"`
    Hash Hash `json:"hash"`
    Data Data `json:"data"`
}

func (b *Block) MarshalJSON() ([]byte, error) {
    err := checkJson(b.data)
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(block{b.header, b.timestamp, b.prevHash, b.hash, b.data})
}

func (b *Block) UnmarshalJSON(data []byte) error {
    _b := block{}
    if err := json.Unmarshal(data, &_b); err != nil {
        return err
    }
    err := checkJson(_b.Data)
    if err != nil {
        return err
    }
    b.header, b.timestamp, b.prevHash, b.hash, b.data = _b.Header, _b.Timestamp, _b.PrevHash, _b.Hash, _b.Data
    return nil
}

func (b *Block) Hash() Hash {
    return b.hash
}

func (b *Block) PrevHash() Hash {
    return b.prevHash
}

func (b *Block) Validate() error {
    if hashS(b.header, b.timestamp.toSting(), b.prevHash.Hex(), b.data.toSting()) == b.hash {
        return nil
    } else {
        return errors.New("hash mismatch")
    }
}

func (b *Block) Print() {
    fmt.Printf("Header: %s; Timestamp: %s; PrevHash: %s...; Hash: %s...;\nData: %s\n",
               b.header, b.timestamp.toTime().String(), b.prevHash.Hex()[:16], b.hash.Hex()[:16], b.data)
}

func (b *Block) Save() ([]byte, error) {
    return json.Marshal(b)
}

func New(
    header string,
    timestamp Time,
    prevHash Hash,
    data Data,
) Block {
    return Block{header, timestamp, prevHash,
                 hashS(header, timestamp.toSting(), prevHash.Hex(), data.toSting()), data}
}

func Load(save []byte) (Block, error) {
    load := Block{}
    err := json.Unmarshal(save, &load)
    return load, err
}

func MBlock(
    header string,
    prevHash Hash,
    data Data,
) Block {
    return New(header, Time(time.Now().Unix()), prevHash, data)
}

func Genesis(data Data) Block {
    return MBlock("Genesis", hashS(), data)
}

func checkJson(inter map[string]interface {}) error {
    for _, value := range inter {
        switch v := value.(type) {
            case bool:
            case string:
            case []string:
            case map[string]interface {}: return checkJson(value.(map[string]interface {}))
            case []map[string]interface {}:
                for _, v := range value.([]map[string]interface {}) {
                    if err := checkJson(v); err != nil {
                        return err
                    }
                    return nil
                }
            default: return errors.New(fmt.Sprint(v, " can't be treated as json object"))
        }
    }
    return nil
}
