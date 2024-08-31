import json
from mido import MidiFile, MidiTrack, Message
import pygame.midi
import pygame

# JSONファイルを読み込む
with open('../notee/output.json', 'r') as file:
    score_data = json.load(file)

# データを抽出
steps = score_data['Starts']
lengths = score_data['Lengths']
notes = score_data['Notes']
length = score_data['Length']

# MIDIファイルとトラックを作成
midi = MidiFile()
track = MidiTrack()
midi.tracks.append(track)

# ノートをMIDIノート番号に変換する関数
def note_to_midi(note):
    return note['Key'] + 12 * note['Octave']

# ノートをトラックに追加
for i, note in enumerate(notes):
    midi_note = note_to_midi(note)
    duration = lengths[i]['Value']  # ここでのdurationはミリ秒として解釈されると仮定しています
    track.append(Message('note_on', note=midi_note, velocity=64, time=0))
    track.append(Message('note_off', note=midi_note, velocity=64, time=duration))

# MIDIファイルを保存
midi.save('output.mid')

# MIDIプレイヤーの初期化
pygame.midi.init()
player = pygame.midi.Output(0)

# MIDIファイルを読み込み、再生
pygame.midi.music.load('output.mid')
pygame.midi.music.play()

# 音楽が再生されるまでスクリプトを実行し続ける
while pygame.midi.music.get_busy():
    pygame.time.wait(100)

# クリーンアップ
pygame.midi.quit()
