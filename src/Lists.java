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

/**
 * List object class.
 */
public class Lists {
    private String listID;
    private String listName;

    /**
     * Constructor which initiates the lists object.
     */
    public Lists(){}

    /**
     * Accessor for the listID variable.
     * @return String   String value representing the ID of the list.
     */
    public String getID(){
        return listID;
    }

    /**
     * Accessor for the listName variable.
     * @return String   String value representing the name of the list.
     */
    public String getName(){
        return listName;
    }

    /**
     * Mutator for the listID variable.
     * @param ID    String value representing the ID of the list.
     */
    public void setID(String ID){
        listID=ID;
    }

    /**
     * Mutator for the listName variable.
     * @param name  String value representing the name of the list.
     */
    public void setName(String name){
        listName=name;
    }
}