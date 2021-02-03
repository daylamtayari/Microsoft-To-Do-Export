/*
 * Copyright (c) 2021 Daylam Tayari <daylam@tayari.gg>
 *
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License version 3as published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with this program.
 * If not see http://www.gnu.org/licenses/ or write to the Free Software Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 *
 * @author Daylam Tayari daylam@tayari.gg https://github.com/daylamtayari
 * @version 1.0
 * Github project home page: https://github.com/daylamtayari/Microsoft-To-Do-Export
 */

import java.util.List;
import java.util.Scanner;

public class Main {
    protected static List<List<String>> contents;
    protected static String token;

    public static void main(String[] args) {
        Scanner sc=new Scanner(System.in);
        System.out.print(
                  "\nWelcome to Microsoft To Do export."
                + "\nAn unofficial tool to export your Microsoft To Do tasks."
                + "\n\nIf you find this tool useful, please consider helping to support me financially:"
                + "\nhttps://paypal.me/daylamtayari https://cash.app/$daylamtayari BTC: 15KcKrsqW6DQdyZPrgRXXmsKkyyZzHAQVX"
                + "\n\nRetrieve your token from..."
                + "\nToken: "
        );
        token=sc.nextLine();
        System.out.print(
                  "\nWhich format do you want to export the tasks:"
                + "\n1. Text format."
                + "\n2. Todoist CSV format."
                + "\n3. Raw JSON format."
        );
        int selection=-1;
        try {
            selection = Integer.parseInt(sc.nextLine());
        }
        catch(Exception ignored){}
        while(selection<1 || selection>3){
            System.out.print("\nINCORRECT INPUT.\nPlease enter a number between 1 and 3: ");
            try{
                selection = Integer.parseInt(sc.nextLine());
            }
            catch(Exception ignored){}
        }
        System.out.print("\nPlease enter the complete file path of where you want to save the output file: ");
        Export.setFP(sc.nextLine());
        System.out.print("\nRetrieving lists...");
        try {
            API.getLists();
        }
        catch(Exception e) {
            System.out.print("\nError retrieving lists, please make sure the token is correct.");
            return;
        }
        System.out.print("\nRetrieved lists.");
        System.out.print("\nLists retrieved: ");
        for(Lists l : API.lists) {
            System.out.print("\n- " + l.getName());
        }
        System.out.print("\nRetrieving tasks...");
        try {
            API.getTasks(selection == 3);
        }
        catch(Exception e) {
            System.out.print("\nError retrieving tasks.");
            return;
        }
        System.out.print("\nRetrieved tasks.");
        System.out.print("\nExporting file...");
        if(selection==1){
            Export.exportText();
        }
        else if(selection==2){
            Export.exportCSV();
        }
        else{
            Export.exportJSON();
        }
        System.out.print(
                  "\n\nFile exported to: "+Export.getFP()+"."
                + "\n\nThank you for using this program."
                + "\n\nIf you find this tool useful, please consider helping to support me financially:"
                + "\nhttps://paypal.me/daylamtayari https://cash.app/$daylamtayari BTC: 15KcKrsqW6DQdyZPrgRXXmsKkyyZzHAQVX"
        );
    }
}