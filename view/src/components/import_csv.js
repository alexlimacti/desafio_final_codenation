import React from "react";

const ImportCsv = () => {
    return(
        <div className="import-csv-form">
            <form>
                <label>
                    URL:
                    <input type="text" name="url"/>
                    <input type="submit" value="browse" />
                    <input type="submit" value="ok" />
                </label>
            </form>
        </div>
    )
}

export default ImportCsv