import json
from mido import MidiFile, MidiTrack, Message

# JSONファイルを読み込む
with open("../main/output.json", "r") as file:
    score_data = json.load(file)

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
    return note["Key"]["Value"] + 12 * note["Octave"]["Value"]


# ノートをトラックに追加
for i, note in enumerate(notes):
    midi_note = note_to_midi(note)
    start = steps[i]["Value"]  # ここでのstartはミリ秒として解釈されると仮定しています
    duration = lengths[i][
        "Value"
    ]  # ここでのdurationはミリ秒として解釈されると仮定しています
    track.append(Message("note_on", note=midi_note, velocity=64, time=start))
    track.append(
        Message("note_off", note=midi_note, velocity=64, time=start + duration)
    )

# MIDIファイルを保存
midi.save("output.mid")
