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

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.ArrayList;
import java.util.List;

import org.apache.http.HttpEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.util.EntityUtils;
import org.json.JSONObject;

/**
 * This method handles all of the
 * API calls and processing.
 */
public class API {
    private static final CloseableHttpClient httpClient=HttpClients.createDefault();
    private static final String API_CORE="https://graph.microsoft.com/v1.0/%s";                 //Formatted string value representing the core of the Microsoft Graph v1.0 API query.
    private static final String LISTS_API="me/todo/lists";                                      //String value representing the list retrieval API suffix.
    private static final String LIST_TASK_API="me/todo/lists/%s/tasks";                         //Formatted string value representing the list tasks retrieval API suffix.
    protected static ArrayList<Lists> lists = new ArrayList<Lists>();                           //Lists arraylist containing the information of the lists.
    protected static List<List<Task>> listContents= new ArrayList<List<Task>>();                //Task arraylist containing the contents of all of the lists.

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
     * @throws IOException
     */
    protected static void getList(String id) throws IOException {
        HttpGet req=new HttpGet(String.format(API_CORE, String.format(LIST_TASK_API, id)));
        req.addHeader("Authorization", Main.token);
        CloseableHttpResponse res=httpClient.execute(req);
        HttpEntity ent= res.getEntity();
        String response= EntityUtils.toString(ent);
        Parser.retrieveContents(response);
    }

    /**
     * This method calls the getList method
     * and retrieves all of the tasks.
     * @throws IOException
     */
    protected static void getTasks() throws IOException {
        for(Lists l: lists){
            getList(l.getID());
        }
    }
}