
package vast

import (
	"fmt"
	"ssp/internal/openrtb"
)

func Build(bid *openrtb.Bid) string {

	return fmt.Sprintf(`
<VAST version="3.0">
 <Ad id="%s">
  <InLine>
   <Creatives>
    <Creative>
     <Linear>
      <MediaFiles>
        <MediaFile type="video/mp4">
        %s
        </MediaFile>
      </MediaFiles>
     </Linear>
    </Creative>
   </Creatives>
  </InLine>
 </Ad>
</VAST>
`, bid.ID, bid.Adm)
}
