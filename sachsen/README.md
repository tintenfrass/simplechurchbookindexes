## Der Trauindex
besteht im Prinzip aus einer Menge von Textdateien worin die Bräutigamme nach Jahr aufgelistet sind, in der Regel in der Reihenfolge wie sie im Kirchenbuch stehen, bei einigen Dateien wurde das Register des Kirchenbuches zur Datenerfasung verwendet.

Durch Shortcuts und Autovervollständigung entspricht die Schreibweise der Namen nicht immer der Schreibweise im Kirchenbuch.<br>
Zusätzlich sind nicht immer alle Namen klar lesbar, heißt der Index enthält eine unbekannte Anzahl von Fehlern.
Die Fehlerquote hängt sehr stark davon ab, wann der Index erstellt wurde und wie gut das Kirchenbuch lesbar ist.
Fehlerkorrekturen werden gerne entgegen genommen.


### Lizenz beachten (Creative Commons Attribution-NonCommercial (CC BY-NC) 3.0)
Die Daten sind für die Forschung gedacht, es ist keine kommerzielle Nutzung erlaubt.


### Online Durchsuchen des Index
https://tintenfrass.github.io/simplechurchbookindexes/

### Wie man möglichst schnell einen einfachen Index erstellt

- man sollte nur die wichtigsten Informationen aufnehmen (Jahr, Namen und fertig)
- bei Trauungen reicht der Bräutigam, wenn man zusätzlich die Braut aus dem Text raussucht braucht man je nach Kirchenbuch ca. 3x solange
- bei Beerdigungen nur die verheirateten Personen aufnehmen spart viel Zeit
- man sollte eine Art Shortcuts/Autovervollständigung für Namen verwenden (siehe unten)
- jedes andere (technische) Hilfsmittel was irgendwie hilft (siehe ganz unten)

### Shortcuts/Autovervollständigung

Welche Art von Shortcuts man benutzt hängt sehr vom Programm ab, mit dem man schreibt und welche Möglichkeiten es bietet.

#### 1. Im Atom-Editor hat sich ein einfaches Snippet bewährt (atom-snippet.txt)
g+tab wird zu Georg, g2+tab zu Gottfried, g3+tab zu Gottlieb usw.
Für die Nachnamen nutzt man die Autovervollständigung, so dass man jeden Nachnamen pro Datei nur einmal schreiben muss.
Damit die Vornamen nicht in der Autovervollständigung für Nachnamen auftauchen, werden diese durch das Snipped mit einem Unterstrich erstellt.
Erst am Ende wenn der Index fertig ist, werden alle Unterstriche entfernt. (Replace-all-Funktion)
Die neueren Index-Listen (ab Sommer 2022) wurden auf diese Weise erstellt.

#### 2. Notepad++ bietet PythonScript-Support (notepad++replace.py)
damit kann man Text automatisch ersetzen lassen z.B: g1 wird zu Georg, g2 zu Gottfried usw.
Das ist ganz gut, kann aber manchmal haken, geht besser.
Einige der älteren Index-Listen wurden auf diese Weise erstellt.

#### 3. Notepad++ hat Tastenkombinationen, die man belegen kann (notepad++shortcuts.xml)
z.B. legt man sich den Namen Georg auf alt+G, Gottfried auf auf alt+str+G und Gottlieb auf alt+shift+G usw.
Das ist ganz ok aber geht besser.
Einige der älteren Index-Listen wurden auf diese Weise erstellt.

#### 4. Speach-To-Text?
Zu langsam und zu anstrengend, dazu sehr hohe Fehlerquote, für viele Nachnamen unbrauchbar, nicht zu empfehlen.

#### 5. ...
Gibt sicher viele weitere Möglichkeiten, muss jeder selber ausprobieren was für ihm am besten geht.

### sonstige Hilfsmittel?

#### mit UniversalSplitscreen kann man den Fokus der Maus auf dem Broserfenser lassen und den Fokus der Tastatur auf dem Editor
Das spart etwas Zeit, weil man nicht bei jedem Scrollen und Umblättern den Fokus wechseln muss,
ist aber etwas tricky einzurichten und klappt nicht mit jeden Editor, z.B. Notepad++ ist dafür schlecht, mit Atom gehts etwas besser.

### Beispiel
wie die Datenerfasung aussehen kann mit dem Atom-Editor mit Shortcuts und Autovervollständigung.<br>
Hier zwei kurze Mitschnitte in Echtzeitgeschwindigkeit für Kirchenbücher, die gut lesbar sind:

https://www.youtube.com/watch?v=MGwKtrvWMSc<br>
https://www.youtube.com/watch?v=aNe8E5PmBKo<br>
(Aus Copyright Gründen wird im Video nur der Editor gezeigt, nicht das Kirchenbuch selber, was im Fenster daneben offen ist.)

Für eine gute Effizienz sollte das Tippen der Namen so wenig Zeit wie möglich benötigen.

Der größte Teil der Zeit geht für das Lesen und Entziffern der Namen im Kirchenbuch drauf, das ist gut so.<br>
Der zweitgrößte Teil geht meist für die Erfassung des Nachnamens drauf, diese lassen sich nicht ganz so gut erfassen wie die zu 99% standartisierten Vornamen.<br>
Der dritte Faktor der Zeit kostet ist Fehlerkorrektur, gerade wenn man schlecht im 10-Finger-Schreiben ist.

Im Idealfall ist es möglich ca. 500 Einträge pro Stunde zu erfassen.
Bei sehr schlecht lesbaren Kirchenbüchern (und/oder sehr exotischen Nachnamen) sind es aber deutlich weniger, ca. 100-200 pro Stunde.