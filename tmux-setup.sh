#!/bin/sh

PWD=$(pwd)
SESSION=$(basename "$PWD")

tmux new-session -s "$SESSION" -d

tmux split-window -v -t "$SESSION"
tmux select-layout even-vertical   # to avoid 'no space for new pane' 
tmux split-window -v -t "$SESSION"
tmux select-layout even-vertical   # to avoid 'no space for new pane' 

tmux send-keys -t "$SESSION:0.0" 'cd client' C-m
tmux send-keys -t "$SESSION:0.1" '(cd server && gow run .)' C-m
tmux send-keys -t "$SESSION:0.2" 'code .' C-m

tmux attach -t "$SESSION"