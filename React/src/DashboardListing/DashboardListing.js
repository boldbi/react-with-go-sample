import React from 'react';
import './DashboardListing.css';
import '../index.css';
import '../index';
import {BoldBI} from '@boldbi/boldbi-embedded-sdk';

//For Bold BI Enterprise edition, it should be like `site/site1`. For Bold BI Cloud, it should be empty string.
const siteIdentifier = "site/site1";

//Bold BI Server URL (ex: http://localhost:5000/bi, http://demo.boldbi.com/bi)
const rootUrl = "http://localhost:5000/bi/";

//Url of the GetDetails action in the Go application(http://localhost:8086/getDetails). Learn more about authorize server [here](https://help.syncfusion.com/bold-bi/embedded-bi/javascript/authorize-server)
const authorizationUrl = "http://localhost:8086/getDetails"

//Your Bold BI application environment. (If Cloud, you should use `cloud`, if Enterprise, you should use `enterprise`)
const environment = "enterprise";

//Set the item id of the dashboard to embed from BI server.Please refer this link(https://help.syncfusion.com/bold-bi/enterprise-bi/share-dashboards/get-dashboard-link#get-link)
const dashboardId = "";
var BoldBiObj;

class DashboardListing extends React.Component {
   constructor(props){
       super(props);
       this.state = {toke: undefined, items: []};
       this.BoldBiObj = new BoldBI();
   };

   renderDashboard() {
    this.dashboard= BoldBI.create({
      serverUrl: rootUrl + siteIdentifier,
      dashboardId: dashboardId,
      embedContainerId: "dashboard",
      embedType: BoldBI.EmbedType.Component,
      environment: environment=="enterprise"? BoldBI.Environment.Enterprise:BoldBI.Environment.Cloud,
      width:"100%",
      height: window.innerHeight + 'px',
      expirationTime:100000,
      authorizationServer: {
          url:authorizationUrl
      }
  });

  console.log(this.dashboard);
  this.dashboard.loadDashboard();     
    
  }

  render() {
    return (
      <div id="DashboardListing">
          <div id="viewer-section">
            <div id="dashboard"></div>
          </div>
      </div>
    );
  }

  componentDidMount() {
    this.renderDashboard();
  }  
}
export default DashboardListing;
