/* tslint:disable */
/* eslint-disable */
/**
 * Myncer API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface OAuthExchangeRequest
 */
export interface OAuthExchangeRequest {
    /**
     * The authorization code returned by the datasource.
     * @type {string}
     * @memberof OAuthExchangeRequest
     */
    code: string;
    /**
     * Optional CSRF protection token returned from the datasource.
     * @type {string}
     * @memberof OAuthExchangeRequest
     */
    state?: string;
}

/**
 * Check if a given object implements the OAuthExchangeRequest interface.
 */
export function instanceOfOAuthExchangeRequest(value: object): value is OAuthExchangeRequest {
    if (!('code' in value) || value['code'] === undefined) return false;
    return true;
}

export function OAuthExchangeRequestFromJSON(json: any): OAuthExchangeRequest {
    return OAuthExchangeRequestFromJSONTyped(json, false);
}

export function OAuthExchangeRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): OAuthExchangeRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'code': json['code'],
        'state': json['state'] == null ? undefined : json['state'],
    };
}

export function OAuthExchangeRequestToJSON(json: any): OAuthExchangeRequest {
    return OAuthExchangeRequestToJSONTyped(json, false);
}

export function OAuthExchangeRequestToJSONTyped(value?: OAuthExchangeRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'code': value['code'],
        'state': value['state'],
    };
}

