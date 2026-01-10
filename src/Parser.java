/*
 * Copyright (c) 2021,2026 Daylam Tayari <daylam@tayari.gg>
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

import org.json.JSONArray;
import org.json.JSONObject;
import java.util.ArrayList;

/**
 * This class is responsible for the
 * parsing of API JSON response.
 */
public class Parser {
    /**
     * This method retrieves the individual
     * lists from the list retrieval API query.
     * @param response      String value representing the total JSON response from the API query.
     */
    protected static void retrieveLists(String response){
        JSONObject jsonResponse=new JSONObject(response);
        JSONArray contents=jsonResponse.getJSONArray("value");
        for(int i=0; i<contents.length(); i++){
            JSONObject jo=contents.getJSONObject(i);
            String name=jo.getString("wellknownListName");
            if(name.equals("none") || name.equals("defaultList")){
               Lists list=new Lists();
               list.setID(jo.getString("id"));
               list.setName(jo.getString("displayName"));
               API.lists.add(list);
            }
        }
    }

    /**
     * This method parses all of the contents (tasks)
     * from a task list retrieval JSON API response.
     * @param response  String value representing the JSON API response contents.
     * @param tasks     ArrayList to add tasks to.
     */
    protected static void retrieveContents(String response, ArrayList<Task> tasks){
        JSONObject jsonResponse=new JSONObject(response);
        JSONArray contents=jsonResponse.getJSONArray("value");
        for(int i=0; i<contents.length(); i++){
            JSONObject jo=contents.getJSONObject(i);
            if(jo.getString("status").equals("notStarted")){
                Task task=new Task();
                task.setTitle(jo.getString("title"));
                if(jo.getString("importance").equals("normal")){
                    task.setImportance(4);
                }
                else{
                    task.setImportance(1);
                }
                try{
                    JSONObject date=jo.getJSONObject("dueDateTime");
                    task.setDate(date.getString("dateTime"));
                    task.setTZ(date.getString("timeZone"));
                }
                catch(Exception e){     //For when a task has no due date.
                    task.setDate("");
                    task.setTZ("");
                }
                try {
                    JSONObject body=jo.getJSONObject("body");
                    String bodyContents=body.getString("content");
                    task.setNote(bodyContents);
                }
                catch(Exception e){
                    task.setNote("");
                }
                tasks.add(task);
            }
        }
    }
}
