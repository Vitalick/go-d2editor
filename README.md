# Diablo 2 Save editor

[![Go Report Card](https://goreportcard.com/badge/github.com/vitalick/go-d2editor)](https://goreportcard.com/report/github.com/vitalick/go-d2editor)
[![GoDoc](https://godoc.org/github.com/vitalick/go-d2editor?status.svg)](https://godoc.org/github.com/vitalick/go-d2editor)

Go package for reading and writing Diablo 2 saves. Supports version 1.10 through Diablo II: Resurrected (1.15).
Supports reading both d2s (player saves) ~~and d2i (shared stash)~~ files.

## Installation

To install the package, use the following:

```bash
go get github.com/vitalick/go-d2editor
```

To install command line program, use the following:

```bash
go install github.com/vitalick/d2editor-cli@latest
```

## Usage

### CLI

For convert JSON to .d2s, use the following:
```bash
d2editor-cli -fromjson <input files>
```

For convert .d2s to JSON, use the following:
```bash
d2editor-cli -tojson <input files>
```

To specify the path to the output folder, we use the following:
```bash
d2editor-cli -fromjson -o <output folder> <input files>
d2editor-cli -tojson -o <output folder> <input files>
```

### As package

Coming soon...


## Roadmap

<details>
<summary>Completed steps</summary>

- ~~Add Character struct~~
- ~~Add Header to Character~~
- ~~Add ActiveWeapon to Character~~
- ~~Add Name to Character~~
- ~~Add Status to Character~~
- ~~Add Progression to Character~~
- ~~Add Unk0x0026 to Character~~
- ~~Add Class to Character~~
- ~~Add Unk0x0029 to Character~~
- ~~Add Level to Character~~
- ~~Add Created to Character~~
- ~~Add LastPlayed to Character~~
- ~~Add Unk0x0034 to Character~~
- ~~Add HotkeySkills to Character~~
- ~~Add LeftSkill to Character~~
- ~~Add RightSkill to Character~~
- ~~Add LeftSwapSkill to Character~~
- ~~Add RightSwapSkill to Character~~
- ~~Add Appearances to Character~~
- ~~Add Locations to Character~~
- ~~Add MapID to Character~~
- ~~Add Unk0x00af to Character~~
- ~~Add Mercenary to Character~~
- ~~Add RealmData to Character~~
- ~~Add Quests to Character~~
- ~~Add Waypoints to Character~~
- ~~Add UnkUnk1 to Character~~
- ~~Add NPCDialogs to Character~~
- ~~Add Attributes to Character~~

</details>

- Add ClassSkills to Character
- Add ItemList to Character
- Add CorpseList to Character
- Add MercenaryItemList to Character
- Add Golem to Character
- Test saved savedata in Diablo 2 Resurrected
- Test saved savedata in Diablo 2 LoD
- Test created savedata in Diablo 2 Resurrected
- Test created savedata in Diablo 2 LoD

## Links

- https://github.com/dschu012/D2SLib (credits to dschu012 for example lib on c#)
- https://github.com/d07RiV/d07riv.github.io/blob/master/d2r.html (credits to d07riv for reversing the item code on D2R)
- https://github.com/nokka/d2s (credits to nokka for structs of classes, skills and items)
- https://github.com/krisives/d2s-format
- http://paul.siramy.free.fr/d2ref/eng/
- http://user.xmission.com/~trevin/DiabloIIv1.09_File_Format.shtml
- https://github.com/nickshanks/Alkor
- https://github.com/HarpyWar/d2s-character-editor


