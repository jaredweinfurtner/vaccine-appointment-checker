# German Vaccination Appointment Checker [(zur deutschen Beschreibung)](README_de.md)

**NOTICE**

As of 26.03.2021 the 116117 vaccination portal has made changes that does not allow this program to function any more (no more REST API access).  Instead, they force people to wait for an indefinite amount of time in front of their computer in a "virtual waiting room".  You have no way of knowing your position in the queue (or even if there is one), so you may be waiting 15 minutes or 6 hours - better not go to the bathroom and miss your spot.  
<br />
<image src="https://user-images.githubusercontent.com/25351659/112604261-37b58600-8e16-11eb-981e-c358d2977a2a.png" width="600"/>

There currently exists a vaccination appointment portal for Germany @ [https://www.impfterminservice](https://www.impfterminservice.de).  However, it has many flaws:

1. You can only search one vaccination center at a time
2. When it searches in the beginning, it searches **all** vaccines, even though they are age restricted
3. If just one of these vaccines is available at the location, it gives you false hope.  When you enter your age, it may say no vaccines are actually available due to age restrictions on the vaccines

## How to use

1. First download the appropriate executable:

    | Operating System | Link    |
    | ---------------- |:-------:|
    | Windows (x86)    | [download](https://github.com/jaredweinfurtner/vaccine-appointment-checker/raw/main/bin/x86/vaccine-appointment-checker.exe) |
    | Windows (x64)    | [download](https://github.com/jaredweinfurtner/vaccine-appointment-checker/raw/main/bin/x64/vaccine-appointment-checker.exe) |
    | Linux            | [download](https://github.com/jaredweinfurtner/vaccine-appointment-checker/raw/main/bin/linux/vaccine-appointment-checker) |

2. Open a terminal and navigate to the download folder
3. List the vaccines (using the downloaded executable) to figure out which one is right for you by looking at the **age** section on each
    ```
    vaccine-appointment-checker.exe -listVaccines
    ```
    This should provide a result similar to:
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
4. Run the scanner by entering your chosen vaccine (by age) and zip codes (comma separated & no spaces)

    ```
   vaccine-appointment-checker.exe -vaccineCode L922 -zipCodes 70174,70376,74081,75175
    ```

    This will result in a list of vaccination centers that have appointments for your chosen vaccine:

    ```
   74081, Nuss??ckerstra??e 3 Heilbronn - Kreisimpfzentrum Heilbronn: 
    https://002-iz.impfterminservice.de/impftermine/service?plz=74081
    
    
    75175, Hohwiesenweg 4 Pforzheim - Kreisimpfzentrum Pforzheim:
    https://005-iz.impfterminservice.de/impftermine/service?plz=75175
    
    Finished searching.

    ```

5. Simply copy-paste the link into your browser to schedule an appointment

**NOTE:** this does not guarantee an appointment since there is no reservation when clicking the link.  If someone schedules the appointment in the time it takes you to fill out the information, then simply go back to step 4 and try again.

## License

Please see [LICENSE](./LICENSE)
