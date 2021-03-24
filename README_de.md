# Deutscher Impfterminprüfer

Aktuell existiert ein Impfterminportal für Deutschland unter [https://www.impfterminservice](https://www.impfterminservice.de). Es hat allerdings einige Schwächen:

1. Man kann nur ein Impfzentrum auf einmal abfragen.
2. Anfangs werden **alle** Impfstoffe durchsucht, selbst die mit Alterseinschränkung.
3. Wenn dort nur ein alterseingeschränkter Impfstoff verfügbar ist, kann das falsche Hoffnunge wecken.

## Benutzung

1. Die passende Anwendung herunterladen :
   
   | Betriebsystem    | Link    |
   | ---------------- |:-------:|
   | Windows (x86)    | [download](https://github.com/jaredweinfurtner/vaccine-appointment-checker/raw/main/bin/x86/vaccine-appointment-checker.exe) |
   | Windows (x64)    | [download](https://github.com/jaredweinfurtner/vaccine-appointment-checker/raw/main/bin/x64/vaccine-appointment-checker.exe) |
   | Linux            | [download](https://github.com/jaredweinfurtner/vaccine-appointment-checker/raw/main/bin/linux/vaccine-appointment-checker) |

2. Eine Kommandozeile öffnen (Unter Windows  [&#8862; Win] + R drücken und dann "cmd" eingeben und bestätigen.)
3. Zum Downloadverzeichnis navigieren (standardmäßig "cd %USERPROFILE%\Downloads")
4. Impfstoffe auflisten (mittels der heruntergeladenen Anwendung).
    ```
    vaccine-appointment-checker.exe -listVaccines
    ```
   Das sollte eine Ausgabe ähnlich der folgenden liefern:
    ```json
    [
        {
            "qualification": "L920",
            "name": "Comirnaty (BioNTech)",
            "interval": 21,
            "age": "16-17"
        },
        {
            "qualification": "L921",
            "name": "mRNA-1273 (Moderna)",
            "tssname": "Moderna & BioNTech",
            "interval": 28,
            "age": "65+"
        },
        {
            "qualification": "L922",
            "name": "COVID-1912 (AstraZeneca)",
            "interval": 63,
            "age": "18-64"
        }
    ]

    ```
5. Um einen passenden Impfstoff zu wählen, soll der Wert von **qualification** gewählt werden, bei dem das eigene Alter von dem Wert **age** abgedeckt wird.
6. Die Anwendung mit dem gewählten Impfstoff ("L922" im Beispiel) und einer kommaseparierten Liste (ohne Leerzeichen) von Postleitzahlen.
    ```
   vaccine-appointment-checker.exe -vaccineCode L922 -zipCodes 70174,70376,74081,75175
    ```
   Das liefert eine Liste von Impfzentren mit verfügbaren Terminen für den gewählten Impfstoff:
    ```
   74081, Nussäckerstraße 3 Heilbronn - Kreisimpfzentrum Heilbronn: 
    https://002-iz.impfterminservice.de/impftermine/service?plz=74081
    
    
    75175, Hohwiesenweg 4 Pforzheim - Kreisimpfzentrum Pforzheim:
    https://005-iz.impfterminservice.de/impftermine/service?plz=75175
    
    Finished searching.

    ```

7. Den Link einfach in die Adresszeile des Webbrowser kopieren um den Termin anzusetzen.

**HINWEIS:** Dies ist keine Garantie für einen Termin, weil das Aufrufen des Links keine Reservierung ist. Es ist möglich, dass jemand anderes den Termin während des Ausfüllens reserviert. In diesem Fall Schritt 6 und folgende wiederholen.

## Lizenz

[LICENSE](./LICENSE) lesen.