# Microsoft To Do Export

### An unofficial solution to export tasks and lists from Microsoft To Do.

## Export Formats:
- Todoist CSV.  
- Human readable text.  
- Raw JSON.  

#### If there is an export format you want me to add, feel free to create an issue suggesting a format.

## Instructions:

1. Download the latest release from the [release tab](https://github.com/daylamtayari/Microsoft-To-Do-Export/releases).  
2. Retrieve your token.  
  a. Go to Microsoft's Graph API Explorer: https://developer.microsoft.com/en-us/graph/graph-explorer  
  b. Sign in with the account that you want to retrieve the tasks from.    
  c. Select the `my To Do task lists` option.    
  d. Click the `Modify permissions` tab and consent to the `Tasks.ReadWrite` permission.   
  e. Navigate to the `Access token` tab and copy the access token.   
     This is the access token you will use, you can now close the Microsoft Graph Explorer.
3. Run the executable.
4. Input your token into the program when prompted.
5. Select your desired output format.
6. Enter your desired output path (complete file path). i.e. `C:\Downloads\export.csv`.

## License:

This project is licensed under GPL v3.0.  
The ocmplete license: [LICENSE](https://github.com/daylamtayari/Microsoft-To-Do-Export/blob/master/LICENSE).  
For more details, please check out the official page: https://www.gnu.org/licenses/gpl-3.0.en.html  

## Support:

If you found this project helpful, please consider checking out my website at [https://tayari.gg](https://tayari.gg) and if you wish help support me financially:  
[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/P5P6AA059)

## Disclaimer:

This project is in no way affiliated with Microsoft and Todoist, and their respective affiliates.
