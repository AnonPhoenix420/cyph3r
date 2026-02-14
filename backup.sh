mkdir -p backups
tar -czvf backups/cyph3r_stable_$(date +%F).tar.gz cmd/ internal/ go.mod
echo "[+] Stable build saved to backups/"
