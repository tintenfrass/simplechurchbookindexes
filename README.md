## Simple Churchbook Indexes
Einfache Indizes von ausgewählten Kirchenbüchern um schneller bestimmte Einträge nachschlagen zu können.

Die vollen Einträge der Kirchenbücher können im sächsischen Landeskirchlichen Archiv oder auf Archion betrachtet werden.



### Wie man möglichst schnell einen einfachen Index erstellt

- man sollte nur die wichtigsten Informationen aufnehmen (Jahr, Namen und fertig)
- bei Trauungen reicht der Bräutigam, wenn man zusätzlich die Braut aus dem Text raussucht braucht man je nach Kirchenbuch ca. 3x solange
- bei Beerdigungen nur die verheirateten Personen aufnehmen spart viel Zeit
- man sollte eine Art Shortcuts/Autovervollständigung für Namen verwenden (siehe unten)
- jedes andere (technische) Hilfsmittel was irgendwie hilft (siehe ganz unten)

### Shortcuts/Autovervollständigung

Welche Art von Shortcuts man benutzt hängt sehr vom Programm ab, mit dem man schreibt und welche Möglichkeiten es bietet.

#### 1. Im Atom-Editor hat sich einfaches Snippet bewährt (atom-snippet.txt)
g+tab wird zu Georg, g2+tab zu Gottfried, g3+tab zu Gottlieb usw.
Für die Nachnamen nutzt man die Autovervollständigung, so dass man jeden Nachnamen pro Datei nur einmal schreiben muss.
Damit die Vornamen nicht in der Autovervollständigung für Nachnamen auftauchen, werden diese durch das Snipped mit einem Unterstrich erstellt.
Erst am Ende wenn der Index fertig ist, werden alle Unterstriche entfernt. (Replace-all-Funktion)
Die neueren Index-Listen wurden auf diese Weise erstellt.

#### 2. Notepad++ bietet PythonScript-Support (notepad++replace.py)
damit kann man Text automatisch ersetzen lassen z.B: g1 wird zu Georg, g2 zu Gottfried usw.
Das ist ganz gut, kann aber manchmal haken, geht besser.
Einige der Index-Listen wurden auf diese Weise erstellt.

#### 3. Notepad++ hat Tastenkombinationen, die man belegen kann (notepad++shortcuts.xml)
z.B. legt man sich den Namen Georg auf alt+G, Gottfried auf auf alt+str+G und Gottlieb auf alt+shift+G usw.
Das ist ganz ok aber geht besser.
Einige der Index-Listen wurden auf diese Weise erstellt.

#### 4. Speach-To-Text?
Zu langsam und zu anstrengend, dazu sehr hohe Fehlerquote, für viele Nachnamen unbrauchbar, nicht zu empfehlen.

#### 5. ...
Gibt sicher viele weitere Möglichkeiten, muss jeder selber ausprobieren was für ihm am besten geht.

### sonstige Hilfsmittel?

#### mit UniversalSplitscreen kann man den Fokus der Maus auf dem Broserfenser lassen und den Fokus der Tastatur auf dem Editor
Das spart etwas Zeit, weil man nicht bei jedem Scrollen und Umblättern den Fokus wechseln muss,
ist aber etwas tricky einzurichten und klappt nicht mit jeden Editor, z.B. Notepad++ ist dafür schlecht, mit Atom gehts etwas besser.
