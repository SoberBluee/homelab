# Output target
Output cron_backup.mp4

# Define the terminal aesthetic
Set Theme "Catppuccin Mocha"
Set FontSize 26          
Set Width 1080
Set Height 1080
Set Padding 30           
Set WindowBar Colorful

# ═══ THE SECRET SETUP (Hidden from the video) ═══
Hide
Type "mkdir -p /tmp/cron_demo && cd /tmp/cron_demo && rm -f crontab"
Enter
Type "crontab() { command nano crontab; }"
Enter
Type "clear"
Enter
Show
# ════════════════════════════════════════════════

# The Typing Animation Sequence (Camera starts here)
Type "# Setting up a Proxmox backup via cronjob"
Sleep 1s
Enter

# This will now visually type out but trigger our clean sandbox file instead
Type "crontab -e"
Sleep 500ms
Enter
Sleep 1.5s

# This will now type cleanly onto line 1 on a fresh screen
Type "  0 3 * * * vzdump 101 --mode snapshot --storage local"
Sleep 1.5s
Enter

# The Perfectly Visual Layout Tree Structure
Type "# │ │ └─── Every Day, Month, and Weekday"
Enter
Type "# │ └───── Hour (3 AM)"
Enter
Type "# └─────── Minute (00)"
Sleep 3s

# Cleanly save and exit Nano so your video loops seamlessly
Ctrl+O
Sleep 500ms
Enter
Ctrl+X
Sleep 1s