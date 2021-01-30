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
 * Object class which handles
 * a task object.
 */
public class Task {
    private String title;       //String value representing the title of the task.
    private int importance;     //Integer value representing the importance level of the task.
    private String dueDate;     //String value representing the due date of the task.
    private String TZ;          //String value representing the timezone of the timestamp.

    /**
     * Constructor for the task object class.
     */
    public Task(){}

    /**
     * Accessor for the title variable.
     * @return String   String variable representing the title of the task.
     */
    public String getTitle(){
        return title;
    }

    /**
     * Accessor for the importance variable.
     * @return int      Integer variable representing the importance of the task.
     */
    public int getImportance(){
        return importance;
    }

    /**
     * Accessor for the due date variable.
     * @return String   String variable representing the due date of the task.
     */
    public String getDate(){
        return dueDate;
    }

    /**
     * Accessor for the timezone variable.
     * @return String   String value representing the timezone of the
     */
    public String getTZ(){
        return TZ;
    }

    /**
     * Mutator for the title variable
     * @param title     String variable representing the title of the task.
     */
    public void setTitle(String title){
        this.title=title;
    }

    /**
     * Mutator for the importance variable.
     * @param importance    Integer variable that represents the importance of the task.
     */
    public void setImportance(int importance){
        this.importance=importance;
    }

    /**
     * Mutator for due date variable.
     * @param date      String variable representing the due date of the task.
     */
    public void setDate(String date){
        dueDate=date;
    }

    /**
     * Mutator for the timezone variable.
     * @param TZ    String variable representing the timezone of the due date.
     */
    public void setTZ(String TZ){
        this.TZ=TZ;
    }
}