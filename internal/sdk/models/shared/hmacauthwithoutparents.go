// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type HMACAuthWithoutParentsConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *HMACAuthWithoutParentsConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type HMACAuthWithoutParents struct {
	Consumer *HMACAuthWithoutParentsConsumer `json:"consumer,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64   `json:"created_at,omitempty"`
	ID        *string  `json:"id,omitempty"`
	Secret    *string  `json:"secret,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	Username  string   `json:"username"`
}

func (o *HMACAuthWithoutParents) GetConsumer() *HMACAuthWithoutParentsConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *HMACAuthWithoutParents) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *HMACAuthWithoutParents) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *HMACAuthWithoutParents) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *HMACAuthWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *HMACAuthWithoutParents) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
