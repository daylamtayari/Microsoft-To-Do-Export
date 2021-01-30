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

import org.json.JSONArray;
import org.json.JSONObject;

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
            if(name.equals("none")){
               API.lists.add(jo);
               API.listIDs.add(jo.getString("id"));
            }
        }
    }
}