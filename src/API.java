/*
 * Copyright (c) 2021, 2025 Daylam Tayari <daylam@tayari.gg>
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

import org.apache.http.HttpEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.util.EntityUtils;
import org.json.JSONObject;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * This method handles all of the
 * API calls and processing.
 */
public class API {
    private static final CloseableHttpClient httpClient=HttpClients.createDefault();
    private static final String API_CORE="https://graph.microsoft.com/v1.0/%s";                 //Formatted string value representing the core of the Microsoft Graph v1.0 API query.
    private static final String LISTS_API="me/todo/lists";                                      //String value representing the list retrieval API suffix.
    private static final String LIST_TASK_API="me/todo/lists/%s/tasks?$skip=%d";                         //Formatted string value representing the list tasks retrieval API suffix.
    protected static ArrayList<Lists> lists = new ArrayList<Lists>();                               //Lists arraylist containing the information of the lists.
    protected static List<List<Task>> listContents= new ArrayList<List<Task>>();                         //Task arraylist containing the contents of all of the lists.
    protected static ArrayList<String> rawJSON=new ArrayList<String>();                         //String arraylist containing raw JSON values of all fo the lists.

    /**
     * This method gets all of the task
     * lists of a user and assigns the
     * values to the listIDs and lists arraylists.
     * @throws IOException
     */
    protected static void getLists() throws IOException {
        HttpGet req=new HttpGet(String.format(API_CORE, LISTS_API));
        req.addHeader("Authorization", Main.token);
        CloseableHttpResponse res=httpClient.execute(req);
        HttpEntity ent= res.getEntity();
        String response= EntityUtils.toString(ent);
        Parser.retrieveLists(response);
    }

    /**
     * This method retrieves the contents of
     * a task list and adds it to the arraylist.
     * @param id    String value representing the ID of the task list.
     * @param json  Boolean value representing whether to store raw JSON.
     * @param skip  Integer value for pagination offset.
     * @param tasks ArrayList to accumulate tasks across pagination, or null for initial call.
     * @throws IOException
     */
    protected static void getList(String id, boolean json, Integer skip, ArrayList<Task> tasks) throws IOException {
        boolean isInitialCall = (tasks == null);
        if(isInitialCall) {
            tasks = new ArrayList<Task>();
        }

        HttpGet req=new HttpGet(String.format(API_CORE, String.format(LIST_TASK_API, id, skip)));
        req.addHeader("Authorization", Main.token);
        CloseableHttpResponse res=httpClient.execute(req);
        HttpEntity ent= res.getEntity();
        String response= EntityUtils.toString(ent);
        JSONObject resJSON=new JSONObject(response);

        // Add opening bracket for this list
        if(json && isInitialCall) {
            rawJSON.add("[");
        }

        if(json){
            String valueString = resJSON.get("value").toString();
            rawJSON.add(valueString.substring(1, valueString.length() - 1));
        }
        else {
            Parser.retrieveContents(response, tasks);
        }

        if (resJSON.has("@odata.nextLink")) {
            String nextLink = resJSON.getString("@odata.nextLink");
            int nextSkip = Integer.parseInt(nextLink.split("\\$skip=")[1]);
            if(json) {
                rawJSON.add(",");
            }
            getList(id, json, nextSkip, tasks);
        }

        // Close this list's array
        if(isInitialCall) {
            if(json) {
                rawJSON.add("],");
            } else {
                listContents.add(tasks);
            }
        }
    }

    /**
     * This method calls the getList method
     * and retrieves all of the tasks.
     * @param json  Boolean value representing whether or not the user's selected output is raw JSON or not.
     * @throws IOException
     */
    protected static void getTasks(boolean json) throws IOException {
        for(Lists l: lists){
            getList(l.getID(), json, 0, null);
        }
        if(json) {
            // Remove trailing comma and add closing bracket
            if(!rawJSON.isEmpty() && rawJSON.get(rawJSON.size() - 1).equals(",")) {
                rawJSON.remove(rawJSON.size() - 1);
            }
        }
    }
}
