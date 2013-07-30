go
==

Go test repository.

Use the following one liner to 'continuously' build and run the app (thanks [0]):
    inotifywait -m -r -e close_write --exclude '.*\.sw[a-z]' src/ | while read line; do make; done

References
----------
[0] http://notes.kel.pe/post/14062831170/continuous-compilation
