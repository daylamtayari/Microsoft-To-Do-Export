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

import java.io.File;
import java.io.FileWriter;
import java.util.ArrayList;

/**
 * This class handles the exporting
 * of the values.
 */
public class Export {
    private static String outputFP;    //String value representing the output total file path.

    /**
     * Method which exports all of the tasks
     * to a CSV file.
     */
    protected static void exportCSV(){
        ArrayList<String> content=new ArrayList<String>();
        content.add("TYPE,CONTENT,PRIORITY,INDENT,AUTHOR,RESPONSIBLE,DATE,DATE_LANG,TIMEZONE");
        for(int i=0; i<API.listContents.size(); i++){
            content.add("section,"+API.lists.get(i).getName()+",,,,,,,");
            for(Task t: API.listContents.get(i)){
                content.add("task,"+t.getTitle()+","+t.getImportance()+",,,,"+t.getDate()+",en,"+t.getTZ());
            }
        }
        write(content);
    }

    /**
     * This method exports all of the tasks
     * to a human readable text file.
     */
    protected static void exportText(){
        ArrayList<String> output=new ArrayList<String>();
        output.add("\tContent\tDue Date\tTimezone");
        for(int i=0; i<API.listContents.size();i++){
            output.add(API.lists.get(i).getName());
            for(Task t: API.listContents.get(i)){
                output.add("\t"+t.getTitle()+"\t"+t.getDate()+"\t"+t.getTZ());
            }
        }
        write(output);
    }

    /**
     * Mutator for the outputFP variable.
     * @param fp    String variable representing the output file path.
     */
    protected static void setFP(String fp){
        outputFP=fp;
    }

    /**
     * This method writes out
     * @param content
     */
    private static void write(ArrayList<String> content){
        File file=new File(outputFP);
        try{
            FileWriter fw=new FileWriter(outputFP);
            for(int i=0; i<content.size(); i++){
                fw.write(content.get(i));
                if(!(i==content.size()-1)){
                    fw.write("\n");
                }
            }
            fw.close();
        }
        catch(Exception ignored){}
    }
}