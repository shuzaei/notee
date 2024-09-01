import json
import sys
from mido import MidiFile, MidiTrack, Message

# JSONファイルを読み込む
# ファイル名を引数で指定
if len(sys.argv) < 3 or not sys.argv[1].endswith(".json") or not sys.argv[2].endswith(".mid"):
    print("Usage: python main.py <filename>.json <output>.mid")
    sys.exit(1)

filename = sys.argv[1]
output = sys.argv[2]

with open(filename) as f:
    score_data = json.load(f)


# データを抽出
steps = score_data["Starts"]
lengths = score_data["Lengths"]
notes = score_data["Notes"]
length = score_data["Length"]["Value"]

# MIDIファイルとトラックを作成
midi = MidiFile()
track = MidiTrack()
midi.tracks.append(track)

# ノートをMIDIノート番号に変換する関数


def note_to_midi(note):
    return note["Key"]["Value"] + 12 * (note["Octave"]["Value"] + 4)


# ノートをトラックに追加
index = 0
lastTime = 0
for time in range(length):
    i = index
    dur = 120  # 1/16
    while i < len(steps) and steps[i]["Value"] == time:
        track.append(
            Message(
                "note_on", note=note_to_midi(notes[i]), time=(time - lastTime) * dur
            )
        )
        lastTime = time + 1
        i += 1
        dur = 0

    i = index
    dur = 120
    while i < len(steps) and steps[i]["Value"] == time:
        track.append(
            Message(
                "note_off",
                note=note_to_midi(notes[i]),
                time=dur,
            )
        )
        i += 1
        dur = 0

    index = i


# MIDIファイルを保存
midi.save(output)
