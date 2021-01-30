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
import org.json.JSONObject;

/**
 * This method handles all of the
 * API calls and processing.
 */
public class API {
    private static final String API_CORE="https://graph.microsoft.com/v1.0/%s";     //Formatted string value representing the core of the Microsoft Graph v1.0 API query.
    private static final String LISTS_API="me/todo/lists";                          //String value representing the list retrieval API suffix.
    private static final String LIST_TASK_API="me/todo/lists/%s/tasks";             //Formatted string value representing the list tasks retrieval API suffix.
    protected static ArrayList<String> listIDs;                                     //String arraylist containing all of the IDs of the task lists of the user.
    //// Its values are parallel to those of the lists and listContents arraylists.
    protected static ArrayList<JSONObject> lists;                                   //JSON object arraylist containing all of the list JSON objects.
    //Its values are parallel to the values of the listIDs and listContents arraylists.
    protected static ArrayList<JSONObject> listContents;                            //JSON object arraylist containing the contents of all of the lists.
    //Its values are parallel to the values of the listIDs and lists arraylists.

}