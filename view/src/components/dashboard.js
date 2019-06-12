import React from "react"
import ImportCsv from "./import_csv"

const Dashboard = () => {
    return(
        <div className="dashboard">
            <h1>Dashboard</h1>
            <div>
                <h2>Número de alertas</h2>
                <h3>##</h3>
            </div>
            <ImportCsv />
        </div>
    )
}

export default Dashboard