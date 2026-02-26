import React from 'react';
import '../index';
import { BoldBI } from '@boldbi/boldbi-embedded-sdk';

//Url of the tokenGeneration action in tokengeneration.go
const tokenGenerationUrl = "http://localhost:8086/tokenGeneration";

class Dashboard extends React.Component {
  constructor(props) {
    super(props);
    this.state = { toke: undefined, items: [] };
    this.BoldBiObj = new BoldBI();
  };

  getEmbedToken() {
        return fetch(tokenGenerationUrl, { // Backend application URL
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({})
        })
          .then(response => {
            if (!response.ok) throw new Error("Token fetch failed");
            return response.text();
          });
      }
    
    renderDashboard(data) {
      this.getEmbedToken()
        .then(accessToken => {
          const dashboard = BoldBI.create({
            serverUrl: data.ServerUrl + "/" + data.SiteIdentifier,
            dashboardId: data.DashboardId,
            embedContainerId: "dashboard",
            embedToken: accessToken
          });

          dashboard.loadDashboard();
        })
        .catch(err => {
          console.error("Error rendering dashboard:", err);
        });
    };

  render() {
    return (
      <div id="DashboardListing">
        <div id="viewer-section">
          <div id="dashboard"></div>
        </div>
      </div>
    );
  }

  async componentDidMount() {
    try {
      const response = await fetch('http://localhost:8086/getServerDetails');
      const embedConfig = await response.json();
      this.renderDashboard(embedConfig);  
    } catch (error) {
      console.log("Error: embedConfig.json file not found.");
      this.setState({ toke: "error", items: "error" });
    }
  }
}
export default Dashboard;